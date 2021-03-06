// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pricing"
)

type Product struct {
	Attributes ProductAttributes `json:"product"`
}

type ProductAttributes struct {
	Attributes EC2Attributes `json:"attributes"`
}

type EC2Attributes struct {
	ServiceCode       string `json:"servicecode"`
	ServiceName       string `json:"servicename"`
	InstanceFamily    string `json:"instanceFamily"`
	InstanceType      string `json:"instanceType"`
	Location          string `json:"location"`
	VCPU              string `json:"vcpu"`
	GPU               string `json:"gpu"`
	Memory            string `json:"memory"`
	PhysicalProcessor string `json:"physicalProcessor"`
	Storage           string `json:"storage"`
}

var packageTemplate = template.Must(template.New("").Parse(`// This file was generated by go generate; DO NOT EDIT
package instancetypes

// RegionTypes returns a list of supported vm types in the region.
func RegionTypes(region string) ([]VM, error) {
	return awsMachines.regionTypes(region)
}

var awsMachines = manager{
	regionVMs: map[string][]VM{
{{- range $k, $v := .RegionInstances }}
		"{{ $k }}": {
{{- range $v }}
			{
				Name:      "{{ .InstanceType }}",
				VCPU:      "{{ .VCPU }}",
				MemoryGiB: "{{ .MemoryGiB }}",
				GPU:       "{{ .GPU }}",
			},
{{- end }}
		},
{{- end }}
	},
}
`))

type instanceType struct {
	InstanceType string
	VCPU         string
	MemoryGiB    string
	GPU          string
}

func main() {
	resolver := endpoints.DefaultResolver()

	regionTypes := make(map[string][]instanceType)
	for _, p := range resolver.(endpoints.EnumPartitions).Partitions() {
		for _, r := range p.Regions() {
			log.Println("Retrieve instance attributes for:", r.ID())
			ec2attrs, err := getEC2types(r.Description())
			handle(err)

			for _, attr := range ec2attrs {
				if regionTypes[r.ID()] == nil {
					regionTypes[r.ID()] = make([]instanceType, 0, len(ec2attrs))
				}
				t := instanceType{
					InstanceType: attr.InstanceType,
					VCPU:         attr.VCPU,
					MemoryGiB:    parseMemory(attr.Memory),
					GPU:          attr.GPU,
				}
				regionTypes[r.ID()] = append(regionTypes[r.ID()], t)
			}
		}
	}

	for region := range regionTypes {
		sort.Slice(regionTypes[region], func(i, j int) bool {
			return regionTypes[region][i].InstanceType < regionTypes[region][j].InstanceType
		})
	}

	f, err := os.Create("vm_types_aws.go")
	if err != nil {
		handle(err)
	}

	defer f.Close()

	err = packageTemplate.Execute(f, struct {
		RegionInstances map[string][]instanceType
	}{
		RegionInstances: regionTypes,
	})

	if err != nil {
		handle(err)
	}
}

func getEC2types(location string) ([]EC2Attributes, error) {
	svc := pricing.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	input := &pricing.GetProductsInput{
		Filters: []*pricing.Filter{
			{
				Field: aws.String("ServiceCode"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("AmazonEC2"),
			},
			{
				Field: aws.String("location"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String(location),
			},
			{
				Field: aws.String("productFamily"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("Compute Instance"),
			},
			{
				Field: aws.String("termType"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("OnDemand"),
			},
			{
				Field: aws.String("operatingSystem"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("Linux"),
			},
			{
				Field: aws.String("operation"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("RunInstances"),
			},
			{
				Field: aws.String("tenancy"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("Shared"),
			},
			{
				Field: aws.String("capacitystatus"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("UnusedCapacityReservation"), // https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html
			},
		},
		FormatVersion: aws.String("aws_v1"),
		MaxResults:    aws.Int64(5),
		ServiceCode:   aws.String("AmazonEC2"),
	}

	priceList := make([]aws.JSONValue, 0)
	for {
		result, err := svc.GetProducts(input)
		if err != nil {
			return nil, err
		}
		priceList = append(priceList, result.PriceList...)
		if result.NextToken == nil {
			break
		}
		input.NextToken = result.NextToken
	}

	ec2types := make([]EC2Attributes, 0)
	for _, product := range priceList {
		if p, ok := product["product"]; ok {
			d, err := json.Marshal(p)
			if err != nil {
				fmt.Println(err)
			}
			info := &ProductAttributes{}
			err = json.Unmarshal(d, info)
			ec2types = append(ec2types, info.Attributes)
		}
	}

	return ec2types, nil
}

func toStrings(in []*pricing.AttributeValue) []string {
	o := make([]string, len(in))
	for i, val := range in {
		if val != nil {
			o[i] = aws.StringValue(val.Value)
		}
	}
	return o
}

func parseMemory(memory string) string {
	reg, err := regexp.Compile("[^0-9\\.]+")
	handle(err)

	return strings.TrimSpace(reg.ReplaceAllString(memory, ""))
}

func handle(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}

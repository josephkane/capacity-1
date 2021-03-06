// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"

	models "github.com/supergiant/capacity/pkg/capacityclient/models"
)

// UpdateConfigReader is a Reader for the UpdateConfig structure.
type UpdateConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateConfigOK creates a UpdateConfigOK with default headers values
func NewUpdateConfigOK() *UpdateConfigOK {
	return &UpdateConfigOK{}
}

/*UpdateConfigOK handles this case with default header values.

configResponse contains an application config parameters.
*/
type UpdateConfigOK struct {
	Payload *models.Config
}

func (o *UpdateConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/config][%d] updateConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Config)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

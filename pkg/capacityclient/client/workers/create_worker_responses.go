// Code generated by go-swagger; DO NOT EDIT.

package workers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/supergiant/capacity/pkg/capacityclient/models"
)

// CreateWorkerReader is a Reader for the CreateWorker structure.
type CreateWorkerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWorkerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateWorkerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateWorkerCreated creates a CreateWorkerCreated with default headers values
func NewCreateWorkerCreated() *CreateWorkerCreated {
	return &CreateWorkerCreated{}
}

/*CreateWorkerCreated handles this case with default header values.

workerResponse contains a worker representation.
*/
type CreateWorkerCreated struct {
	Payload *models.Worker
}

func (o *CreateWorkerCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/workers][%d] createWorkerCreated  %+v", 201, o.Payload)
}

func (o *CreateWorkerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Worker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
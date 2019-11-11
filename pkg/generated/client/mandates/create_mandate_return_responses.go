// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// CreateMandateReturnReader is a Reader for the CreateMandateReturn structure.
type CreateMandateReturnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMandateReturnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateMandateReturnCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateMandateReturnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateMandateReturnCreated creates a CreateMandateReturnCreated with default headers values
func NewCreateMandateReturnCreated() *CreateMandateReturnCreated {
	return &CreateMandateReturnCreated{}
}

/*CreateMandateReturnCreated handles this case with default header values.

Return creation response
*/
type CreateMandateReturnCreated struct {

	//Payload

	// isStream: false
	*models.MandateReturnCreationResponse
}

func (o *CreateMandateReturnCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/mandates/{id}/returns][%d] createMandateReturnCreated  %+v", 201, o)
}

func (o *CreateMandateReturnCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateReturnCreationResponse = new(models.MandateReturnCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateReturnCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMandateReturnBadRequest creates a CreateMandateReturnBadRequest with default headers values
func NewCreateMandateReturnBadRequest() *CreateMandateReturnBadRequest {
	return &CreateMandateReturnBadRequest{}
}

/*CreateMandateReturnBadRequest handles this case with default header values.

Return creation error
*/
type CreateMandateReturnBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateMandateReturnBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/mandates/{id}/returns][%d] createMandateReturnBadRequest  %+v", 400, o)
}

func (o *CreateMandateReturnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// CreatePaymentRecallDecisionReader is a Reader for the CreatePaymentRecallDecision structure.
type CreatePaymentRecallDecisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentRecallDecisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentRecallDecisionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentRecallDecisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentRecallDecisionCreated creates a CreatePaymentRecallDecisionCreated with default headers values
func NewCreatePaymentRecallDecisionCreated() *CreatePaymentRecallDecisionCreated {
	return &CreatePaymentRecallDecisionCreated{}
}

/*CreatePaymentRecallDecisionCreated handles this case with default header values.

Recall decision creation response
*/
type CreatePaymentRecallDecisionCreated struct {

	//Payload

	// isStream: false
	*models.RecallDecisionCreationResponse
}

func (o *CreatePaymentRecallDecisionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/recalls/{recallId}/decisions][%d] createPaymentRecallDecisionCreated", 201)
}

func (o *CreatePaymentRecallDecisionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallDecisionCreationResponse = new(models.RecallDecisionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallDecisionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRecallDecisionBadRequest creates a CreatePaymentRecallDecisionBadRequest with default headers values
func NewCreatePaymentRecallDecisionBadRequest() *CreatePaymentRecallDecisionBadRequest {
	return &CreatePaymentRecallDecisionBadRequest{}
}

/*CreatePaymentRecallDecisionBadRequest handles this case with default header values.

Recall decision creation error
*/
type CreatePaymentRecallDecisionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreatePaymentRecallDecisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/recalls/{recallId}/decisions][%d] createPaymentRecallDecisionBadRequest", 400)
}

func (o *CreatePaymentRecallDecisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// CreateClaimReversalReader is a Reader for the CreateClaimReversal structure.
type CreateClaimReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClaimReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateClaimReversalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateClaimReversalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateClaimReversalCreated creates a CreateClaimReversalCreated with default headers values
func NewCreateClaimReversalCreated() *CreateClaimReversalCreated {
	return &CreateClaimReversalCreated{}
}

/*CreateClaimReversalCreated handles this case with default header values.

Claim Reversal creation response
*/
type CreateClaimReversalCreated struct {

	//Payload

	// isStream: false
	*models.ClaimReversalDetailsResponse
}

func (o *CreateClaimReversalCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/claims/{id}/reversals][%d] createClaimReversalCreated", 201)
}

func (o *CreateClaimReversalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimReversalDetailsResponse = new(models.ClaimReversalDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ClaimReversalDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClaimReversalBadRequest creates a CreateClaimReversalBadRequest with default headers values
func NewCreateClaimReversalBadRequest() *CreateClaimReversalBadRequest {
	return &CreateClaimReversalBadRequest{}
}

/*CreateClaimReversalBadRequest handles this case with default header values.

Claim Reversal creation error
*/
type CreateClaimReversalBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateClaimReversalBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/claims/{id}/reversals][%d] createClaimReversalBadRequest", 400)
}

func (o *CreateClaimReversalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

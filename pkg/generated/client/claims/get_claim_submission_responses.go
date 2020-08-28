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

// GetClaimSubmissionReader is a Reader for the GetClaimSubmission structure.
type GetClaimSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClaimSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClaimSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetClaimSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClaimSubmissionOK creates a GetClaimSubmissionOK with default headers values
func NewGetClaimSubmissionOK() *GetClaimSubmissionOK {
	return &GetClaimSubmissionOK{}
}

/*GetClaimSubmissionOK handles this case with default header values.

Claim Submission details
*/
type GetClaimSubmissionOK struct {

	//Payload

	// isStream: false
	*models.ClaimSubmissionDetailsResponse
}

func (o *GetClaimSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/submissions/{submissionId}][%d] getClaimSubmissionOK", 200)
}

func (o *GetClaimSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimSubmissionDetailsResponse = new(models.ClaimSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ClaimSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClaimSubmissionBadRequest creates a GetClaimSubmissionBadRequest with default headers values
func NewGetClaimSubmissionBadRequest() *GetClaimSubmissionBadRequest {
	return &GetClaimSubmissionBadRequest{}
}

/*GetClaimSubmissionBadRequest handles this case with default header values.

Error
*/
type GetClaimSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetClaimSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/submissions/{submissionId}][%d] getClaimSubmissionBadRequest", 400)
}

func (o *GetClaimSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

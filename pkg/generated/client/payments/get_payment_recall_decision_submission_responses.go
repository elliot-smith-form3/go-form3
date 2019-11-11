// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetPaymentRecallDecisionSubmissionReader is a Reader for the GetPaymentRecallDecisionSubmission structure.
type GetPaymentRecallDecisionSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRecallDecisionSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentRecallDecisionSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRecallDecisionSubmissionOK creates a GetPaymentRecallDecisionSubmissionOK with default headers values
func NewGetPaymentRecallDecisionSubmissionOK() *GetPaymentRecallDecisionSubmissionOK {
	return &GetPaymentRecallDecisionSubmissionOK{}
}

/*GetPaymentRecallDecisionSubmissionOK handles this case with default header values.

Recall decision submission details
*/
type GetPaymentRecallDecisionSubmissionOK struct {

	//Payload

	// isStream: false
	*models.RecallDecisionSubmissionDetailsResponse
}

func (o *GetPaymentRecallDecisionSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/submissions/{submissionId}][%d] getPaymentRecallDecisionSubmissionOK  %+v", 200, o)
}

func (o *GetPaymentRecallDecisionSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallDecisionSubmissionDetailsResponse = new(models.RecallDecisionSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallDecisionSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

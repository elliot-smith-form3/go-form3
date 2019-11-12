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

// GetPaymentRecallDecisionReader is a Reader for the GetPaymentRecallDecision structure.
type GetPaymentRecallDecisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRecallDecisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentRecallDecisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRecallDecisionOK creates a GetPaymentRecallDecisionOK with default headers values
func NewGetPaymentRecallDecisionOK() *GetPaymentRecallDecisionOK {
	return &GetPaymentRecallDecisionOK{}
}

/*GetPaymentRecallDecisionOK handles this case with default header values.

Recall decision details
*/
type GetPaymentRecallDecisionOK struct {

	//Payload

	// isStream: false
	*models.RecallDecisionDetailsResponse
}

func (o *GetPaymentRecallDecisionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}][%d] getPaymentRecallDecisionOK  %+v", 200, o)
}

func (o *GetPaymentRecallDecisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallDecisionDetailsResponse = new(models.RecallDecisionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallDecisionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

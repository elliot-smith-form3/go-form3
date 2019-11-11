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

// GetPaymentReturnReversalAdmissionReader is a Reader for the GetPaymentReturnReversalAdmission structure.
type GetPaymentReturnReversalAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReturnReversalAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentReturnReversalAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReturnReversalAdmissionOK creates a GetPaymentReturnReversalAdmissionOK with default headers values
func NewGetPaymentReturnReversalAdmissionOK() *GetPaymentReturnReversalAdmissionOK {
	return &GetPaymentReturnReversalAdmissionOK{}
}

/*GetPaymentReturnReversalAdmissionOK handles this case with default header values.

Return reversal admission details
*/
type GetPaymentReturnReversalAdmissionOK struct {

	//Payload

	// isStream: false
	*models.ReturnReversalAdmissionDetailsResponse
}

func (o *GetPaymentReturnReversalAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/returns/{returnId}/reversals/{reversalId}/admissions/{admissionId}][%d] getPaymentReturnReversalAdmissionOK  %+v", 200, o)
}

func (o *GetPaymentReturnReversalAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnReversalAdmissionDetailsResponse = new(models.ReturnReversalAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReturnReversalAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetPaymentReturnReader is a Reader for the GetPaymentReturn structure.
type GetPaymentReturnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReturnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentReturnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReturnOK creates a GetPaymentReturnOK with default headers values
func NewGetPaymentReturnOK() *GetPaymentReturnOK {
	return &GetPaymentReturnOK{}
}

/*GetPaymentReturnOK handles this case with default header values.

Return details
*/
type GetPaymentReturnOK struct {

	//Payload

	// isStream: false
	*models.ReturnDetailsResponse
}

func (o *GetPaymentReturnOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/returns/{returnId}][%d] getPaymentReturnOK  %+v", 200, o)
}

func (o *GetPaymentReturnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnDetailsResponse = new(models.ReturnDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReturnDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

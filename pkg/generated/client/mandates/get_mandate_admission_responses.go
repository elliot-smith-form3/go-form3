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

// GetMandateAdmissionReader is a Reader for the GetMandateAdmission structure.
type GetMandateAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMandateAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMandateAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMandateAdmissionOK creates a GetMandateAdmissionOK with default headers values
func NewGetMandateAdmissionOK() *GetMandateAdmissionOK {
	return &GetMandateAdmissionOK{}
}

/*GetMandateAdmissionOK handles this case with default header values.

Mandate Admission details
*/
type GetMandateAdmissionOK struct {

	//Payload

	// isStream: false
	*models.MandateAdmissionDetailsResponse
}

func (o *GetMandateAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates/{id}/admissions/{admissionId}][%d] getMandateAdmissionOK  %+v", 200, o)
}

func (o *GetMandateAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateAdmissionDetailsResponse = new(models.MandateAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

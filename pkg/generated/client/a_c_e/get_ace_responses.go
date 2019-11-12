// Code generated by go-swagger; DO NOT EDIT.

package a_c_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetAceReader is a Reader for the GetAce structure.
type GetAceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAceOK creates a GetAceOK with default headers values
func NewGetAceOK() *GetAceOK {
	return &GetAceOK{}
}

/*GetAceOK handles this case with default header values.

ACE details
*/
type GetAceOK struct {

	//Payload

	// isStream: false
	*models.AceDetailsResponse
}

func (o *GetAceOK) Error() string {
	return fmt.Sprintf("[GET /security/roles/{role_id}/aces/{ace_id}][%d] getAceOK  %+v", 200, o)
}

func (o *GetAceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AceDetailsResponse = new(models.AceDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AceDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

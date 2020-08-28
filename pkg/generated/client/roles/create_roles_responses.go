// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// CreateRolesReader is a Reader for the CreateRoles structure.
type CreateRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateRolesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRolesCreated creates a CreateRolesCreated with default headers values
func NewCreateRolesCreated() *CreateRolesCreated {
	return &CreateRolesCreated{}
}

/*CreateRolesCreated handles this case with default header values.

Role creation response
*/
type CreateRolesCreated struct {

	//Payload

	// isStream: false
	*models.RoleCreationResponse
}

func (o *CreateRolesCreated) Error() string {
	return fmt.Sprintf("[POST /security/roles][%d] createRolesCreated", 201)
}

func (o *CreateRolesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RoleCreationResponse = new(models.RoleCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RoleCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

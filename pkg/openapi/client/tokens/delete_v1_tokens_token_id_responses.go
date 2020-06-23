// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1TokensTokenIDReader is a Reader for the DeleteV1TokensTokenID structure.
type DeleteV1TokensTokenIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1TokensTokenIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1TokensTokenIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteV1TokensTokenIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1TokensTokenIDOK creates a DeleteV1TokensTokenIDOK with default headers values
func NewDeleteV1TokensTokenIDOK() *DeleteV1TokensTokenIDOK {
	return &DeleteV1TokensTokenIDOK{}
}

/*DeleteV1TokensTokenIDOK handles this case with default header values.

Successful Response
*/
type DeleteV1TokensTokenIDOK struct {
}

func (o *DeleteV1TokensTokenIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/tokens/{tokenID}][%d] deleteV1TokensTokenIdOK ", 200)
}

func (o *DeleteV1TokensTokenIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1TokensTokenIDNotFound creates a DeleteV1TokensTokenIDNotFound with default headers values
func NewDeleteV1TokensTokenIDNotFound() *DeleteV1TokensTokenIDNotFound {
	return &DeleteV1TokensTokenIDNotFound{}
}

/*DeleteV1TokensTokenIDNotFound handles this case with default header values.

Token not found
*/
type DeleteV1TokensTokenIDNotFound struct {
}

func (o *DeleteV1TokensTokenIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/tokens/{tokenID}][%d] deleteV1TokensTokenIdNotFound ", 404)
}

func (o *DeleteV1TokensTokenIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

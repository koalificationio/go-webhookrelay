// Code generated by go-swagger; DO NOT EDIT.

package functions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1FunctionsFunctionIDConfigConfigKeyReader is a Reader for the DeleteV1FunctionsFunctionIDConfigConfigKey structure.
type DeleteV1FunctionsFunctionIDConfigConfigKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1FunctionsFunctionIDConfigConfigKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1FunctionsFunctionIDConfigConfigKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteV1FunctionsFunctionIDConfigConfigKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1FunctionsFunctionIDConfigConfigKeyOK creates a DeleteV1FunctionsFunctionIDConfigConfigKeyOK with default headers values
func NewDeleteV1FunctionsFunctionIDConfigConfigKeyOK() *DeleteV1FunctionsFunctionIDConfigConfigKeyOK {
	return &DeleteV1FunctionsFunctionIDConfigConfigKeyOK{}
}

/*DeleteV1FunctionsFunctionIDConfigConfigKeyOK handles this case with default header values.

Successful Response
*/
type DeleteV1FunctionsFunctionIDConfigConfigKeyOK struct {
}

func (o *DeleteV1FunctionsFunctionIDConfigConfigKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/functions/{functionID}/config/{configKey}][%d] deleteV1FunctionsFunctionIdConfigConfigKeyOK ", 200)
}

func (o *DeleteV1FunctionsFunctionIDConfigConfigKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1FunctionsFunctionIDConfigConfigKeyNotFound creates a DeleteV1FunctionsFunctionIDConfigConfigKeyNotFound with default headers values
func NewDeleteV1FunctionsFunctionIDConfigConfigKeyNotFound() *DeleteV1FunctionsFunctionIDConfigConfigKeyNotFound {
	return &DeleteV1FunctionsFunctionIDConfigConfigKeyNotFound{}
}

/*DeleteV1FunctionsFunctionIDConfigConfigKeyNotFound handles this case with default header values.

Function config value not found
*/
type DeleteV1FunctionsFunctionIDConfigConfigKeyNotFound struct {
}

func (o *DeleteV1FunctionsFunctionIDConfigConfigKeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/functions/{functionID}/config/{configKey}][%d] deleteV1FunctionsFunctionIdConfigConfigKeyNotFound ", 404)
}

func (o *DeleteV1FunctionsFunctionIDConfigConfigKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

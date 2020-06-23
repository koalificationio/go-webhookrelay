// Code generated by go-swagger; DO NOT EDIT.

package tunnels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

// GetV1TunnelsReader is a Reader for the GetV1Tunnels structure.
type GetV1TunnelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TunnelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TunnelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1TunnelsOK creates a GetV1TunnelsOK with default headers values
func NewGetV1TunnelsOK() *GetV1TunnelsOK {
	return &GetV1TunnelsOK{}
}

/*GetV1TunnelsOK handles this case with default header values.

Successful Response
*/
type GetV1TunnelsOK struct {
	Payload []*models.Tunnel
}

func (o *GetV1TunnelsOK) Error() string {
	return fmt.Sprintf("[GET /v1/tunnels][%d] getV1TunnelsOK  %+v", 200, o.Payload)
}

func (o *GetV1TunnelsOK) GetPayload() []*models.Tunnel {
	return o.Payload
}

func (o *GetV1TunnelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

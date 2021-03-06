// Code generated by go-swagger; DO NOT EDIT.

package tunnels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tunnels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tunnels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1TunnelsTunnelID(params *DeleteV1TunnelsTunnelIDParams) (*DeleteV1TunnelsTunnelIDOK, error)

	GetV1Tunnels(params *GetV1TunnelsParams) (*GetV1TunnelsOK, error)

	GetV1TunnelsTunnelID(params *GetV1TunnelsTunnelIDParams) (*GetV1TunnelsTunnelIDOK, error)

	PostV1Tunnels(params *PostV1TunnelsParams) (*PostV1TunnelsCreated, error)

	PutV1TunnelsTunnelID(params *PutV1TunnelsTunnelIDParams) (*PutV1TunnelsTunnelIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteV1TunnelsTunnelID deletes tunnel
*/
func (a *Client) DeleteV1TunnelsTunnelID(params *DeleteV1TunnelsTunnelIDParams) (*DeleteV1TunnelsTunnelIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1TunnelsTunnelIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteV1TunnelsTunnelID",
		Method:             "DELETE",
		PathPattern:        "/v1/tunnels/{tunnelID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1TunnelsTunnelIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1TunnelsTunnelIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteV1TunnelsTunnelID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1Tunnels lists all tunnels
*/
func (a *Client) GetV1Tunnels(params *GetV1TunnelsParams) (*GetV1TunnelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TunnelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV1Tunnels",
		Method:             "GET",
		PathPattern:        "/v1/tunnels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TunnelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TunnelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Tunnels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1TunnelsTunnelID gets tunnel details
*/
func (a *Client) GetV1TunnelsTunnelID(params *GetV1TunnelsTunnelIDParams) (*GetV1TunnelsTunnelIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TunnelsTunnelIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV1TunnelsTunnelID",
		Method:             "GET",
		PathPattern:        "/v1/tunnels/{tunnelID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TunnelsTunnelIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TunnelsTunnelIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1TunnelsTunnelID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1Tunnels creates a new tunnel

  You may create your own tunnel using this action. It takes a JSON object containing a tunnel request with a specified destination. Paid plans have an option to specify either encryption, subdomain or full domain.
*/
func (a *Client) PostV1Tunnels(params *PostV1TunnelsParams) (*PostV1TunnelsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1TunnelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostV1Tunnels",
		Method:             "POST",
		PathPattern:        "/v1/tunnels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1TunnelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1TunnelsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1Tunnels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutV1TunnelsTunnelID updates tunnel details
*/
func (a *Client) PutV1TunnelsTunnelID(params *PutV1TunnelsTunnelIDParams) (*PutV1TunnelsTunnelIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutV1TunnelsTunnelIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutV1TunnelsTunnelID",
		Method:             "PUT",
		PathPattern:        "/v1/tunnels/{tunnelID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutV1TunnelsTunnelIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutV1TunnelsTunnelIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutV1TunnelsTunnelID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddAmazonRDSMySQLService adds amazon RDS my SQL service adds amazon RDS my SQL service
*/
func (a *Client) AddAmazonRDSMySQLService(params *AddAmazonRDSMySQLServiceParams) (*AddAmazonRDSMySQLServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddAmazonRDSMySQLServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddAmazonRDSMySQLService",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Services/AddAmazonRDSMySQL",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddAmazonRDSMySQLServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddAmazonRDSMySQLServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddAmazonRDSMySQLServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddMongoDBService adds mongo DB service adds mongo DB service
*/
func (a *Client) AddMongoDBService(params *AddMongoDBServiceParams) (*AddMongoDBServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMongoDBServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMongoDBService",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Services/AddMongoDB",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMongoDBServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddMongoDBServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddMongoDBServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddMySQLService adds my SQL service adds my SQL service
*/
func (a *Client) AddMySQLService(params *AddMySQLServiceParams) (*AddMySQLServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMySQLServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMySQLService",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Services/AddMySQL",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMySQLServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddMySQLServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddMySQLServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddPostgreSQLService adds postgre SQL service adds postgre SQL service
*/
func (a *Client) AddPostgreSQLService(params *AddPostgreSQLServiceParams) (*AddPostgreSQLServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPostgreSQLServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddPostgreSQLService",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Services/AddPostgreSQL",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPostgreSQLServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddPostgreSQLServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddPostgreSQLServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddProxySQLService adds proxy SQL service adds proxy SQL service
*/
func (a *Client) AddProxySQLService(params *AddProxySQLServiceParams) (*AddProxySQLServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddProxySQLServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddProxySQLService",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Services/AddProxySQL",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddProxySQLServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddProxySQLServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddProxySQLServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetService gets service returns a single service by ID
*/
func (a *Client) GetService(params *GetServiceParams) (*GetServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetService",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Services/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListServices lists services returns a list of all services
*/
func (a *Client) ListServices(params *ListServicesParams) (*ListServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListServices",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Services/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListServicesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RemoveService removes service removes service
*/
func (a *Client) RemoveService(params *RemoveServiceParams) (*RemoveServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveService",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Services/Remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

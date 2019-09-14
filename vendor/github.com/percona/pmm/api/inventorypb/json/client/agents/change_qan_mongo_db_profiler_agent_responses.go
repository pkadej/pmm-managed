// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeQANMongoDBProfilerAgentReader is a Reader for the ChangeQANMongoDBProfilerAgent structure.
type ChangeQANMongoDBProfilerAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeQANMongoDBProfilerAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeQANMongoDBProfilerAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeQANMongoDBProfilerAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeQANMongoDBProfilerAgentOK creates a ChangeQANMongoDBProfilerAgentOK with default headers values
func NewChangeQANMongoDBProfilerAgentOK() *ChangeQANMongoDBProfilerAgentOK {
	return &ChangeQANMongoDBProfilerAgentOK{}
}

/*ChangeQANMongoDBProfilerAgentOK handles this case with default header values.

A successful response.
*/
type ChangeQANMongoDBProfilerAgentOK struct {
	Payload *ChangeQANMongoDBProfilerAgentOKBody
}

func (o *ChangeQANMongoDBProfilerAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANMongoDBProfilerAgent][%d] changeQanMongoDbProfilerAgentOk  %+v", 200, o.Payload)
}

func (o *ChangeQANMongoDBProfilerAgentOK) GetPayload() *ChangeQANMongoDBProfilerAgentOKBody {
	return o.Payload
}

func (o *ChangeQANMongoDBProfilerAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeQANMongoDBProfilerAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeQANMongoDBProfilerAgentDefault creates a ChangeQANMongoDBProfilerAgentDefault with default headers values
func NewChangeQANMongoDBProfilerAgentDefault(code int) *ChangeQANMongoDBProfilerAgentDefault {
	return &ChangeQANMongoDBProfilerAgentDefault{
		_statusCode: code,
	}
}

/*ChangeQANMongoDBProfilerAgentDefault handles this case with default header values.

An error response.
*/
type ChangeQANMongoDBProfilerAgentDefault struct {
	_statusCode int

	Payload *ChangeQANMongoDBProfilerAgentDefaultBody
}

// Code gets the status code for the change QAN mongo DB profiler agent default response
func (o *ChangeQANMongoDBProfilerAgentDefault) Code() int {
	return o._statusCode
}

func (o *ChangeQANMongoDBProfilerAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANMongoDBProfilerAgent][%d] ChangeQANMongoDBProfilerAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeQANMongoDBProfilerAgentDefault) GetPayload() *ChangeQANMongoDBProfilerAgentDefaultBody {
	return o.Payload
}

func (o *ChangeQANMongoDBProfilerAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeQANMongoDBProfilerAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeQANMongoDBProfilerAgentBody change QAN mongo DB profiler agent body
swagger:model ChangeQANMongoDBProfilerAgentBody
*/
type ChangeQANMongoDBProfilerAgentBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeQANMongoDBProfilerAgentParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change QAN mongo DB profiler agent body
func (o *ChangeQANMongoDBProfilerAgentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMongoDBProfilerAgentBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMongoDBProfilerAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANMongoDBProfilerAgentDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangeQANMongoDBProfilerAgentDefaultBody
*/
type ChangeQANMongoDBProfilerAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change QAN mongo DB profiler agent default body
func (o *ChangeQANMongoDBProfilerAgentDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMongoDBProfilerAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANMongoDBProfilerAgentOKBody change QAN mongo DB profiler agent OK body
swagger:model ChangeQANMongoDBProfilerAgentOKBody
*/
type ChangeQANMongoDBProfilerAgentOKBody struct {

	// qan mongodb profiler agent
	QANMongodbProfilerAgent *ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent `json:"qan_mongodb_profiler_agent,omitempty"`
}

// Validate validates this change QAN mongo DB profiler agent OK body
func (o *ChangeQANMongoDBProfilerAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANMongodbProfilerAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMongoDBProfilerAgentOKBody) validateQANMongodbProfilerAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMongodbProfilerAgent) { // not required
		return nil
	}

	if o.QANMongodbProfilerAgent != nil {
		if err := o.QANMongodbProfilerAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQanMongoDbProfilerAgentOk" + "." + "qan_mongodb_profiler_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMongoDBProfilerAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent QANMongoDBProfilerAgent runs within pmm-agent and sends MongoDB Query Analytics data to the PMM Server.
swagger:model ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent
*/
type ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if query examples are disabled.
	// bool query_examples_disabled = 8; TODO https://jira.percona.com/browse/PMM-4650
	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// MongoDB password for getting profiler data.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// MongoDB username for getting profiler data.
	Username string `json:"username,omitempty"`
}

// Validate validates this change QAN mongo DB profiler agent OK body QAN mongodb profiler agent
func (o *ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeQanMongoDbProfilerAgentOkBodyQanMongodbProfilerAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeQanMongoDbProfilerAgentOkBodyQanMongodbProfilerAgentTypeStatusPropEnum = append(changeQanMongoDbProfilerAgentOkBodyQanMongodbProfilerAgentTypeStatusPropEnum, v)
	}
}

const (

	// ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusSTARTING captures enum value "STARTING"
	ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusSTARTING string = "STARTING"

	// ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusRUNNING captures enum value "RUNNING"
	ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusRUNNING string = "RUNNING"

	// ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusWAITING captures enum value "WAITING"
	ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusWAITING string = "WAITING"

	// ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusSTOPPING captures enum value "STOPPING"
	ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusSTOPPING string = "STOPPING"

	// ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusDONE captures enum value "DONE"
	ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusDONE string = "DONE"
)

// prop value enum
func (o *ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, changeQanMongoDbProfilerAgentOkBodyQanMongodbProfilerAgentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeQanMongoDbProfilerAgentOk"+"."+"qan_mongodb_profiler_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) UnmarshalBinary(b []byte) error {
	var res ChangeQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANMongoDBProfilerAgentParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeQANMongoDBProfilerAgentParamsBodyCommon
*/
type ChangeQANMongoDBProfilerAgentParamsBodyCommon struct {

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`
}

// Validate validates this change QAN mongo DB profiler agent params body common
func (o *ChangeQANMongoDBProfilerAgentParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMongoDBProfilerAgentParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeQANMongoDBProfilerAgentParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

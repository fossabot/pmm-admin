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

// ChangeMySQLdExporterReader is a Reader for the ChangeMySQLdExporter structure.
type ChangeMySQLdExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeMySQLdExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeMySQLdExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewChangeMySQLdExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeMySQLdExporterOK creates a ChangeMySQLdExporterOK with default headers values
func NewChangeMySQLdExporterOK() *ChangeMySQLdExporterOK {
	return &ChangeMySQLdExporterOK{}
}

/*ChangeMySQLdExporterOK handles this case with default header values.

A successful response.
*/
type ChangeMySQLdExporterOK struct {
	Payload *ChangeMySQLdExporterOKBody
}

func (o *ChangeMySQLdExporterOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/ChangeMySQLdExporter][%d] changeMySQLdExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeMySQLdExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeMySQLdExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeMySQLdExporterDefault creates a ChangeMySQLdExporterDefault with default headers values
func NewChangeMySQLdExporterDefault(code int) *ChangeMySQLdExporterDefault {
	return &ChangeMySQLdExporterDefault{
		_statusCode: code,
	}
}

/*ChangeMySQLdExporterDefault handles this case with default header values.

An error response.
*/
type ChangeMySQLdExporterDefault struct {
	_statusCode int

	Payload *ChangeMySQLdExporterDefaultBody
}

// Code gets the status code for the change my s q ld exporter default response
func (o *ChangeMySQLdExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeMySQLdExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/ChangeMySQLdExporter][%d] ChangeMySQLdExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeMySQLdExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeMySQLdExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeMySQLdExporterBody change my s q ld exporter body
swagger:model ChangeMySQLdExporterBody
*/
type ChangeMySQLdExporterBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeMySQLdExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change my s q ld exporter body
func (o *ChangeMySQLdExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMySQLdExporterBody) validateCommon(formats strfmt.Registry) error {

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
func (o *ChangeMySQLdExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMySQLdExporterDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangeMySQLdExporterDefaultBody
*/
type ChangeMySQLdExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change my s q ld exporter default body
func (o *ChangeMySQLdExporterDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMySQLdExporterOKBody change my s q ld exporter OK body
swagger:model ChangeMySQLdExporterOKBody
*/
type ChangeMySQLdExporterOKBody struct {

	// mysqld exporter
	MysqldExporter *ChangeMySQLdExporterOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`
}

// Validate validates this change my s q ld exporter OK body
func (o *ChangeMySQLdExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMySQLdExporterOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeMySQLdExporterOk" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMySQLdExporterOKBodyMysqldExporter MySQLdExporter runs on Generic or Container Node and exposes MySQL and AmazonRDSMySQL Service metrics.
swagger:model ChangeMySQLdExporterOKBodyMysqldExporter
*/
type ChangeMySQLdExporterOKBodyMysqldExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this change my s q ld exporter OK body mysqld exporter
func (o *ChangeMySQLdExporterOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum = append(changeMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusSTARTING captures enum value "STARTING"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusSTARTING string = "STARTING"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusRUNNING captures enum value "RUNNING"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusRUNNING string = "RUNNING"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusWAITING captures enum value "WAITING"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusWAITING string = "WAITING"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusSTOPPING string = "STOPPING"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusDONE captures enum value "DONE"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *ChangeMySQLdExporterOKBodyMysqldExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, changeMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ChangeMySQLdExporterOKBodyMysqldExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeMySQLdExporterOk"+"."+"mysqld_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMySQLdExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeMySQLdExporterParamsBodyCommon
*/
type ChangeMySQLdExporterParamsBodyCommon struct {

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`
}

// Validate validates this change my s q ld exporter params body common
func (o *ChangeMySQLdExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

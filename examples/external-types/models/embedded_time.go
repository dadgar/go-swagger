// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	timeext "time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EmbeddedTime
// This type demonstrates how we can embed an external type and wraps the validation.
//
// This is especially useful if you want to reuse some types from the standard library,
// such as `time.Time` or `json.RawMessage`.
//
//
// swagger:model EmbeddedTime
type EmbeddedTime struct {
	timeext.Time
}

func (m EmbeddedTime) Validate(formats strfmt.Registry) error {
	var f interface{} = m.Time
	if v, ok := f.(runtime.Validatable); ok {
		return v.Validate(formats)
	}
	return nil
}

func (m EmbeddedTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var f interface{} = m.Time
	if v, ok := f.(runtime.ContextValidatable); ok {
		return v.ContextValidate(ctx, formats)
	}
	return nil
}
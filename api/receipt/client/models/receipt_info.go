// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReceiptInfo Receipt info
// swagger:model ReceiptInfo
type ReceiptInfo struct {

	// Full text
	FullText string `json:"fullText,omitempty"`

	// Receipt lang
	Lang string `json:"lang,omitempty"`

	// Receipt title
	Title string `json:"title,omitempty"`

	// TotalPrice
	TotalPrice float64 `json:"totalPrice,omitempty"`
}

// Validate validates this receipt info
func (m *ReceiptInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ReceiptInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReceiptInfo) UnmarshalBinary(b []byte) error {
	var res ReceiptInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIOAuth2FederationUpdateOAuthServiceBodyChanges Configuration changes.
//
// Configuration changes.
//
// swagger:model apiOAuth2FederationUpdateOAuthServiceBodyChanges
type APIOAuth2FederationUpdateOAuthServiceBodyChanges struct {

	// OAuth 2.0 Authorization Endpoint
	AuthURL string `json:"auth_url,omitempty"`

	// Identity claims policy rules
	// NOTE: Order matters
	//
	// google.protobuf.Struct claims = 7;
	Claims []*APIClaim `json:"claims"`

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// unix
	CreatedAt string `json:"created_at,omitempty"`

	// user
	CreatedBy *APIObjectID `json:"created_by,omitempty"`

	// OpenID Connect Service Discovery
	DiscoveryURL string `json:"discovery_url,omitempty"`

	// domain
	Domain *APIObjectID `json:"domain,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// logo
	Logo string `json:"logo,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// display
	Name string `json:"name,omitempty"`

	// Scopes to be requested
	Scopes []string `json:"scopes"`

	// OAuth 2.0 Token Endpoint
	TokenURL string `json:"token_url,omitempty"`

	// well-known vendor; provider
	Type string `json:"type,omitempty"`

	// unix
	UpdatedAt string `json:"updated_at,omitempty"`

	// user
	UpdatedBy *APIObjectID `json:"updated_by,omitempty"`

	// OpenID Connect Userinfo Endpoint
	UserinfoURL string `json:"userinfo_url,omitempty"`
}

// Validate validates this api o auth2 federation update o auth service body changes
func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaims(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) validateClaims(formats strfmt.Registry) error {
	if swag.IsZero(m.Claims) { // not required
		return nil
	}

	for i := 0; i < len(m.Claims); i++ {
		if swag.IsZero(m.Claims[i]) { // not required
			continue
		}

		if m.Claims[i] != nil {
			if err := m.Claims[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claims" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("claims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) validateUpdatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this api o auth2 federation update o auth service body changes based on the context it is used
func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClaims(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) contextValidateClaims(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Claims); i++ {

		if m.Claims[i] != nil {

			if swag.IsZero(m.Claims[i]) { // not required
				return nil
			}

			if err := m.Claims[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claims" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("claims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) contextValidateDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.Domain != nil {

		if swag.IsZero(m.Domain) { // not required
			return nil
		}

		if err := m.Domain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedBy != nil {

		if swag.IsZero(m.UpdatedBy) { // not required
			return nil
		}

		if err := m.UpdatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIOAuth2FederationUpdateOAuthServiceBodyChanges) UnmarshalBinary(b []byte) error {
	var res APIOAuth2FederationUpdateOAuthServiceBodyChanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

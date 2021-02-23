// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BookSlotOfFacilityHandlerFunc turns a function with the right signature into a book slot of facility handler
type BookSlotOfFacilityHandlerFunc func(BookSlotOfFacilityParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookSlotOfFacilityHandlerFunc) Handle(params BookSlotOfFacilityParams) middleware.Responder {
	return fn(params)
}

// BookSlotOfFacilityHandler interface for that can handle valid book slot of facility params
type BookSlotOfFacilityHandler interface {
	Handle(BookSlotOfFacilityParams) middleware.Responder
}

// NewBookSlotOfFacility creates a new http.Handler for the book slot of facility operation
func NewBookSlotOfFacility(ctx *middleware.Context, handler BookSlotOfFacilityHandler) *BookSlotOfFacility {
	return &BookSlotOfFacility{Context: ctx, Handler: handler}
}

/*BookSlotOfFacility swagger:route POST /appointment bookSlotOfFacility

Book a slot in facility

*/
type BookSlotOfFacility struct {
	Context *middleware.Context
	Handler BookSlotOfFacilityHandler
}

func (o *BookSlotOfFacility) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookSlotOfFacilityParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// BookSlotOfFacilityBody book slot of facility body
//
// swagger:model BookSlotOfFacilityBody
type BookSlotOfFacilityBody struct {

	// enrollment code
	// Required: true
	EnrollmentCode *string `json:"enrollmentCode"`

	// facility slot Id
	// Required: true
	FacilitySlotID *string `json:"facilitySlotId"`
}

// Validate validates this book slot of facility body
func (o *BookSlotOfFacilityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnrollmentCode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFacilitySlotID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BookSlotOfFacilityBody) validateEnrollmentCode(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"enrollmentCode", "body", o.EnrollmentCode); err != nil {
		return err
	}

	return nil
}

func (o *BookSlotOfFacilityBody) validateFacilitySlotID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"facilitySlotId", "body", o.FacilitySlotID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BookSlotOfFacilityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BookSlotOfFacilityBody) UnmarshalBinary(b []byte) error {
	var res BookSlotOfFacilityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

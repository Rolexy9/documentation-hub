// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetContractStoreParams creates a new GetContractStoreParams object
// with the default values initialized.
func NewGetContractStoreParams() *GetContractStoreParams {
	var ()
	return &GetContractStoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContractStoreParamsWithTimeout creates a new GetContractStoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContractStoreParamsWithTimeout(timeout time.Duration) *GetContractStoreParams {
	var ()
	return &GetContractStoreParams{

		timeout: timeout,
	}
}

// NewGetContractStoreParamsWithContext creates a new GetContractStoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContractStoreParamsWithContext(ctx context.Context) *GetContractStoreParams {
	var ()
	return &GetContractStoreParams{

		Context: ctx,
	}
}

// NewGetContractStoreParamsWithHTTPClient creates a new GetContractStoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContractStoreParamsWithHTTPClient(client *http.Client) *GetContractStoreParams {
	var ()
	return &GetContractStoreParams{
		HTTPClient: client,
	}
}

/*GetContractStoreParams contains all the parameters to send to the API endpoint
for the get contract store operation typically these are written to a http.Request
*/
type GetContractStoreParams struct {

	/*Pubkey
	  The pubkey of the contract

	*/
	Pubkey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get contract store params
func (o *GetContractStoreParams) WithTimeout(timeout time.Duration) *GetContractStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contract store params
func (o *GetContractStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contract store params
func (o *GetContractStoreParams) WithContext(ctx context.Context) *GetContractStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contract store params
func (o *GetContractStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contract store params
func (o *GetContractStoreParams) WithHTTPClient(client *http.Client) *GetContractStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contract store params
func (o *GetContractStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPubkey adds the pubkey to the get contract store params
func (o *GetContractStoreParams) WithPubkey(pubkey string) *GetContractStoreParams {
	o.SetPubkey(pubkey)
	return o
}

// SetPubkey adds the pubkey to the get contract store params
func (o *GetContractStoreParams) SetPubkey(pubkey string) {
	o.Pubkey = pubkey
}

// WriteToRequest writes these params to a swagger request
func (o *GetContractStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pubkey
	if err := r.SetPathParam("pubkey", o.Pubkey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Rekor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

package entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetLogEntryByUUIDHandlerFunc turns a function with the right signature into a get log entry by UUID handler
type GetLogEntryByUUIDHandlerFunc func(GetLogEntryByUUIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLogEntryByUUIDHandlerFunc) Handle(params GetLogEntryByUUIDParams) middleware.Responder {
	return fn(params)
}

// GetLogEntryByUUIDHandler interface for that can handle valid get log entry by UUID params
type GetLogEntryByUUIDHandler interface {
	Handle(GetLogEntryByUUIDParams) middleware.Responder
}

// NewGetLogEntryByUUID creates a new http.Handler for the get log entry by UUID operation
func NewGetLogEntryByUUID(ctx *middleware.Context, handler GetLogEntryByUUIDHandler) *GetLogEntryByUUID {
	return &GetLogEntryByUUID{Context: ctx, Handler: handler}
}

/*GetLogEntryByUUID swagger:route GET /api/v1/log/entries/{entryUUID} entries getLogEntryByUuid

Retrieves an entry from the transparency log (if it exists) by UUID

*/
type GetLogEntryByUUID struct {
	Context *middleware.Context
	Handler GetLogEntryByUUIDHandler
}

func (o *GetLogEntryByUUID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLogEntryByUUIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
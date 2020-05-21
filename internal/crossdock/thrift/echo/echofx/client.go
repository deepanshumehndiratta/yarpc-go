// Code generated by thriftrw-plugin-yarpc
// @generated

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package echofx

import (
	fx "go.uber.org/fx"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	restriction "go.uber.org/yarpc/api/x/restriction"
	thrift "go.uber.org/yarpc/encoding/thrift"
	echoclient "go.uber.org/yarpc/internal/crossdock/thrift/echo/echoclient"
)

// Params defines the dependencies for the Echo client.
type Params struct {
	fx.In

	Provider    yarpc.ClientConfig
	Restriction restriction.Checker `optional:"true"`
}

// Result defines the output of the Echo client module. It provides a
// Echo client to an Fx application.
type Result struct {
	fx.Out

	Client echoclient.Interface

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// Client provides a Echo client to an Fx application using the given name
// for routing.
//
// 	fx.Provide(
// 		echofx.Client("..."),
// 		newHandler,
// 	)
func Client(name string, opts ...thrift.ClientOption) interface{} {
	return func(p Params) Result {
		cc := p.Provider.ClientConfig(name)
		if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok && p.Restriction != nil {
			if err := p.Restriction.Check(thrift.Encoding, namer.TransportName()); err != nil {
				panic(err.Error())
			}
		}
		client := echoclient.New(cc, opts...)
		return Result{Client: client}
	}
}

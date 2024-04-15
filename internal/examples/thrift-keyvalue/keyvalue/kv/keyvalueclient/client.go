// Code generated by thriftrw-plugin-yarpc
// @generated

// Copyright (c) 2024 Uber Technologies, Inc.
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

package keyvalueclient

import (
	context "context"
	wire "go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	kv "go.uber.org/yarpc/internal/examples/thrift-keyvalue/keyvalue/kv"
	reflect "reflect"
)

// Interface is a client for the KeyValue service.
type Interface interface {
	GetValue(
		ctx context.Context,
		Key *string,
		opts ...yarpc.CallOption,
	) (string, error)

	SetValue(
		ctx context.Context,
		Key *string,
		Value *string,
		opts ...yarpc.CallOption,
	) error
}

// New builds a new client for the KeyValue service.
//
//	client := keyvalueclient.New(dispatcher.ClientConfig("keyvalue"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "KeyValue",
			ClientConfig: c,
		}, opts...),
		nwc: thrift.NewNoWire(thrift.Config{
			Service:      "KeyValue",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	c   thrift.Client
	nwc thrift.NoWireClient
}

func (c client) GetValue(
	ctx context.Context,
	_Key *string,
	opts ...yarpc.CallOption,
) (success string, err error) {

	var result kv.KeyValue_GetValue_Result
	args := kv.KeyValue_GetValue_Helper.Args(_Key)

	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = kv.KeyValue_GetValue_Helper.UnwrapResponse(&result)
	return
}

func (c client) SetValue(
	ctx context.Context,
	_Key *string,
	_Value *string,
	opts ...yarpc.CallOption,
) (err error) {

	var result kv.KeyValue_SetValue_Result
	args := kv.KeyValue_SetValue_Helper.Args(_Key, _Value)

	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	err = kv.KeyValue_SetValue_Helper.UnwrapResponse(&result)
	return
}

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: encoding/protobuf/internal/testpb/test.proto

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

package testpb

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// TestYARPCClient is the YARPC client-side interface for the Test service.
type TestYARPCClient interface {
	Unary(context.Context, *TestMessage, ...yarpc.CallOption) (*TestMessage, error)
	Duplex(context.Context, ...yarpc.CallOption) (TestServiceDuplexYARPCClient, error)
}

// TestServiceDuplexYARPCClient sends TestMessages and receives TestMessages, returning io.EOF when the stream is complete.
type TestServiceDuplexYARPCClient interface {
	Context() context.Context
	Send(*TestMessage, ...yarpc.StreamOption) error
	Recv(...yarpc.StreamOption) (*TestMessage, error)
	CloseSend(...yarpc.StreamOption) error
}

func newTestYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) TestYARPCClient {
	return &_TestYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.encoding.protobuf.Test",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewTestYARPCClient builds a new YARPC client for the Test service.
func NewTestYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) TestYARPCClient {
	return newTestYARPCClient(clientConfig, nil, options...)
}

// TestYARPCServer is the YARPC server-side interface for the Test service.
type TestYARPCServer interface {
	Unary(context.Context, *TestMessage) (*TestMessage, error)
	Duplex(TestServiceDuplexYARPCServer) error
}

// TestServiceDuplexYARPCServer receives TestMessages and sends TestMessage.
type TestServiceDuplexYARPCServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*TestMessage, error)
	Send(*TestMessage, ...yarpc.StreamOption) error
}

type buildTestYARPCProceduresParams struct {
	Server      TestYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildTestYARPCProcedures(params buildTestYARPCProceduresParams) []transport.Procedure {
	handler := &_TestYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.yarpc.encoding.protobuf.Test",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Unary",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Unary,
							NewRequest:  newTestServiceUnaryYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{
				{
					MethodName: "Duplex",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.Duplex,
						},
					),
				},
			},
		},
	)
}

// BuildTestYARPCProcedures prepares an implementation of the Test service for YARPC registration.
func BuildTestYARPCProcedures(server TestYARPCServer) []transport.Procedure {
	return buildTestYARPCProcedures(buildTestYARPCProceduresParams{Server: server})
}

// FxTestYARPCClientParams defines the input
// for NewFxTestYARPCClient. It provides the
// paramaters to get a TestYARPCClient in an
// Fx application.
type FxTestYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxTestYARPCClientResult defines the output
// of NewFxTestYARPCClient. It provides a
// TestYARPCClient to an Fx application.
type FxTestYARPCClientResult struct {
	fx.Out

	Client TestYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxTestYARPCClient provides a TestYARPCClient
// to an Fx application using the given name for routing.
//
//	fx.Provide(
//	  testpb.NewFxTestYARPCClient("service-name"),
//	  ...
//	)
func NewFxTestYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxTestYARPCClientParams) FxTestYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxTestYARPCClientResult{
			Client: newTestYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxTestYARPCProceduresParams defines the input
// for NewFxTestYARPCProcedures. It provides the
// paramaters to get TestYARPCServer procedures in an
// Fx application.
type FxTestYARPCProceduresParams struct {
	fx.In

	Server      TestYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxTestYARPCProceduresResult defines the output
// of NewFxTestYARPCProcedures. It provides
// TestYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxTestYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxTestYARPCProcedures provides TestYARPCServer procedures to an Fx application.
// It expects a TestYARPCServer to be present in the container.
//
//	fx.Provide(
//	  testpb.NewFxTestYARPCProcedures(),
//	  ...
//	)
func NewFxTestYARPCProcedures() interface{} {
	return func(params FxTestYARPCProceduresParams) FxTestYARPCProceduresResult {
		return FxTestYARPCProceduresResult{
			Procedures: buildTestYARPCProcedures(buildTestYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: TestReflectionMeta,
		}
	}
}

// TestReflectionMeta is the reflection server metadata
// required for using the gRPC reflection protocol with YARPC.
//
// See https://github.com/grpc/grpc/blob/master/doc/server-reflection.md.
var TestReflectionMeta = reflection.ServerMeta{
	ServiceName:     "uber.yarpc.encoding.protobuf.Test",
	FileDescriptors: yarpcFileDescriptorClosurefc320162ebaf2b25,
}

type _TestYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_TestYARPCCaller) Unary(ctx context.Context, request *TestMessage, options ...yarpc.CallOption) (*TestMessage, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Unary", request, newTestServiceUnaryYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*TestMessage)
	if !ok {
		return nil, protobuf.CastError(emptyTestServiceUnaryYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_TestYARPCCaller) Duplex(ctx context.Context, options ...yarpc.CallOption) (TestServiceDuplexYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "Duplex", options...)
	if err != nil {
		return nil, err
	}
	return &_TestServiceDuplexYARPCClient{stream: stream}, nil
}

type _TestYARPCHandler struct {
	server TestYARPCServer
}

func (h *_TestYARPCHandler) Unary(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *TestMessage
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*TestMessage)
		if !ok {
			return nil, protobuf.CastError(emptyTestServiceUnaryYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Unary(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_TestYARPCHandler) Duplex(serverStream *protobuf.ServerStream) error {
	return h.server.Duplex(&_TestServiceDuplexYARPCServer{serverStream: serverStream})
}

type _TestServiceDuplexYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_TestServiceDuplexYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_TestServiceDuplexYARPCClient) Send(request *TestMessage, options ...yarpc.StreamOption) error {
	return c.stream.Send(request, options...)
}

func (c *_TestServiceDuplexYARPCClient) Recv(options ...yarpc.StreamOption) (*TestMessage, error) {
	responseMessage, err := c.stream.Receive(newTestServiceDuplexYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*TestMessage)
	if !ok {
		return nil, protobuf.CastError(emptyTestServiceDuplexYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_TestServiceDuplexYARPCClient) CloseSend(options ...yarpc.StreamOption) error {
	return c.stream.Close(options...)
}

type _TestServiceDuplexYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_TestServiceDuplexYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_TestServiceDuplexYARPCServer) Recv(options ...yarpc.StreamOption) (*TestMessage, error) {
	requestMessage, err := s.serverStream.Receive(newTestServiceDuplexYARPCRequest, options...)
	if requestMessage == nil {
		return nil, err
	}
	request, ok := requestMessage.(*TestMessage)
	if !ok {
		return nil, protobuf.CastError(emptyTestServiceDuplexYARPCRequest, requestMessage)
	}
	return request, err
}

func (s *_TestServiceDuplexYARPCServer) Send(response *TestMessage, options ...yarpc.StreamOption) error {
	return s.serverStream.Send(response, options...)
}

func newTestServiceUnaryYARPCRequest() proto.Message {
	return &TestMessage{}
}

func newTestServiceUnaryYARPCResponse() proto.Message {
	return &TestMessage{}
}

func newTestServiceDuplexYARPCRequest() proto.Message {
	return &TestMessage{}
}

func newTestServiceDuplexYARPCResponse() proto.Message {
	return &TestMessage{}
}

var (
	emptyTestServiceUnaryYARPCRequest   = &TestMessage{}
	emptyTestServiceUnaryYARPCResponse  = &TestMessage{}
	emptyTestServiceDuplexYARPCRequest  = &TestMessage{}
	emptyTestServiceDuplexYARPCResponse = &TestMessage{}
)

var yarpcFileDescriptorClosurefc320162ebaf2b25 = [][]byte{
	// encoding/protobuf/internal/testpb/test.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xcd, 0x4b, 0xce,
		0x4f, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0xcc,
		0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x49, 0x2d, 0x2e, 0x29, 0x48, 0x02, 0x53, 0x7a,
		0x60, 0x59, 0x21, 0x99, 0xd2, 0xa4, 0xd4, 0x22, 0xbd, 0xca, 0xc4, 0xa2, 0x82, 0x64, 0x3d, 0x98,
		0x46, 0x3d, 0x98, 0x46, 0x25, 0x65, 0x2e, 0xee, 0x90, 0xd4, 0xe2, 0x12, 0xdf, 0xd4, 0xe2, 0xe2,
		0xc4, 0xf4, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46,
		0x0d, 0xce, 0x20, 0x08, 0xc7, 0xe8, 0x24, 0x23, 0x17, 0x0b, 0x48, 0x95, 0x50, 0x2c, 0x17, 0x6b,
		0x68, 0x5e, 0x62, 0x51, 0xa5, 0x90, 0xa6, 0x1e, 0x3e, 0x53, 0xf5, 0x90, 0x8c, 0x94, 0x22, 0x5e,
		0xa9, 0x50, 0x12, 0x17, 0x9b, 0x4b, 0x69, 0x41, 0x4e, 0x6a, 0x05, 0x6d, 0xcc, 0xd7, 0x60, 0x34,
		0x60, 0x74, 0xe2, 0x88, 0x62, 0x83, 0x84, 0x51, 0x12, 0x1b, 0x58, 0x8d, 0x31, 0x20, 0x00, 0x00,
		0xff, 0xff, 0x2c, 0x33, 0x41, 0xcb, 0x4f, 0x01, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) TestYARPCClient {
			return NewTestYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}

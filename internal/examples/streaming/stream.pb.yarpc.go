// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: internal/examples/streaming/stream.proto

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

package streaming

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

// HelloYARPCClient is the YARPC client-side interface for the Hello service.
type HelloYARPCClient interface {
	HelloUnary(context.Context, *HelloRequest, ...yarpc.CallOption) (*HelloResponse, error)
	HelloOutStream(context.Context, ...yarpc.CallOption) (HelloServiceHelloOutStreamYARPCClient, error)
	HelloInStream(context.Context, *HelloRequest, ...yarpc.CallOption) (HelloServiceHelloInStreamYARPCClient, error)
	HelloThere(context.Context, ...yarpc.CallOption) (HelloServiceHelloThereYARPCClient, error)
}

// HelloServiceHelloOutStreamYARPCClient sends HelloRequests and receives the single HelloResponse when sending is done.
type HelloServiceHelloOutStreamYARPCClient interface {
	Context() context.Context
	Send(*HelloRequest, ...yarpc.StreamOption) error
	CloseAndRecv(...yarpc.StreamOption) (*HelloResponse, error)
}

// HelloServiceHelloInStreamYARPCClient receives HelloResponses, returning io.EOF when the stream is complete.
type HelloServiceHelloInStreamYARPCClient interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*HelloResponse, error)
	CloseSend(...yarpc.StreamOption) error
}

// HelloServiceHelloThereYARPCClient sends HelloRequests and receives HelloResponses, returning io.EOF when the stream is complete.
type HelloServiceHelloThereYARPCClient interface {
	Context() context.Context
	Send(*HelloRequest, ...yarpc.StreamOption) error
	Recv(...yarpc.StreamOption) (*HelloResponse, error)
	CloseSend(...yarpc.StreamOption) error
}

func newHelloYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) HelloYARPCClient {
	return &_HelloYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.streaming.Hello",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewHelloYARPCClient builds a new YARPC client for the Hello service.
func NewHelloYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) HelloYARPCClient {
	return newHelloYARPCClient(clientConfig, nil, options...)
}

// HelloYARPCServer is the YARPC server-side interface for the Hello service.
type HelloYARPCServer interface {
	HelloUnary(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloOutStream(HelloServiceHelloOutStreamYARPCServer) (*HelloResponse, error)
	HelloInStream(*HelloRequest, HelloServiceHelloInStreamYARPCServer) error
	HelloThere(HelloServiceHelloThereYARPCServer) error
}

// HelloServiceHelloOutStreamYARPCServer receives HelloRequests.
type HelloServiceHelloOutStreamYARPCServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*HelloRequest, error)
}

// HelloServiceHelloInStreamYARPCServer sends HelloResponses.
type HelloServiceHelloInStreamYARPCServer interface {
	Context() context.Context
	Send(*HelloResponse, ...yarpc.StreamOption) error
}

// HelloServiceHelloThereYARPCServer receives HelloRequests and sends HelloResponse.
type HelloServiceHelloThereYARPCServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*HelloRequest, error)
	Send(*HelloResponse, ...yarpc.StreamOption) error
}

type buildHelloYARPCProceduresParams struct {
	Server      HelloYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildHelloYARPCProcedures(params buildHelloYARPCProceduresParams) []transport.Procedure {
	handler := &_HelloYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.yarpc.internal.examples.streaming.Hello",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "HelloUnary",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.HelloUnary,
							NewRequest:  newHelloServiceHelloUnaryYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{
				{
					MethodName: "HelloThere",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.HelloThere,
						},
					),
				},

				{
					MethodName: "HelloInStream",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.HelloInStream,
						},
					),
				},

				{
					MethodName: "HelloOutStream",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.HelloOutStream,
						},
					),
				},
			},
		},
	)
}

// BuildHelloYARPCProcedures prepares an implementation of the Hello service for YARPC registration.
func BuildHelloYARPCProcedures(server HelloYARPCServer) []transport.Procedure {
	return buildHelloYARPCProcedures(buildHelloYARPCProceduresParams{Server: server})
}

// FxHelloYARPCClientParams defines the input
// for NewFxHelloYARPCClient. It provides the
// paramaters to get a HelloYARPCClient in an
// Fx application.
type FxHelloYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxHelloYARPCClientResult defines the output
// of NewFxHelloYARPCClient. It provides a
// HelloYARPCClient to an Fx application.
type FxHelloYARPCClientResult struct {
	fx.Out

	Client HelloYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxHelloYARPCClient provides a HelloYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    streaming.NewFxHelloYARPCClient("service-name"),
//    ...
//  )
func NewFxHelloYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxHelloYARPCClientParams) FxHelloYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxHelloYARPCClientResult{
			Client: newHelloYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxHelloYARPCProceduresParams defines the input
// for NewFxHelloYARPCProcedures. It provides the
// paramaters to get HelloYARPCServer procedures in an
// Fx application.
type FxHelloYARPCProceduresParams struct {
	fx.In

	Server      HelloYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxHelloYARPCProceduresResult defines the output
// of NewFxHelloYARPCProcedures. It provides
// HelloYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxHelloYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxHelloYARPCProcedures provides HelloYARPCServer procedures to an Fx application.
// It expects a HelloYARPCServer to be present in the container.
//
//  fx.Provide(
//    streaming.NewFxHelloYARPCProcedures(),
//    ...
//  )
func NewFxHelloYARPCProcedures() interface{} {
	return func(params FxHelloYARPCProceduresParams) FxHelloYARPCProceduresResult {
		return FxHelloYARPCProceduresResult{
			Procedures: buildHelloYARPCProcedures(buildHelloYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "uber.yarpc.internal.examples.streaming.Hello",
				FileDescriptors: yarpcFileDescriptorClosure45d12c3ddf34baf8,
			},
		}
	}
}

type _HelloYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_HelloYARPCCaller) HelloUnary(ctx context.Context, request *HelloRequest, options ...yarpc.CallOption) (*HelloResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "HelloUnary", request, newHelloServiceHelloUnaryYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HelloResponse)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloUnaryYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_HelloYARPCCaller) HelloOutStream(ctx context.Context, options ...yarpc.CallOption) (HelloServiceHelloOutStreamYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "HelloOutStream", options...)
	if err != nil {
		return nil, err
	}
	return &_HelloServiceHelloOutStreamYARPCClient{stream: stream}, nil
}

func (c *_HelloYARPCCaller) HelloInStream(ctx context.Context, request *HelloRequest, options ...yarpc.CallOption) (HelloServiceHelloInStreamYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "HelloInStream", options...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(request); err != nil {
		return nil, err
	}
	return &_HelloServiceHelloInStreamYARPCClient{stream: stream}, nil
}

func (c *_HelloYARPCCaller) HelloThere(ctx context.Context, options ...yarpc.CallOption) (HelloServiceHelloThereYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "HelloThere", options...)
	if err != nil {
		return nil, err
	}
	return &_HelloServiceHelloThereYARPCClient{stream: stream}, nil
}

type _HelloYARPCHandler struct {
	server HelloYARPCServer
}

func (h *_HelloYARPCHandler) HelloUnary(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *HelloRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*HelloRequest)
		if !ok {
			return nil, protobuf.CastError(emptyHelloServiceHelloUnaryYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.HelloUnary(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_HelloYARPCHandler) HelloOutStream(serverStream *protobuf.ServerStream) error {
	response, err := h.server.HelloOutStream(&_HelloServiceHelloOutStreamYARPCServer{serverStream: serverStream})
	if err != nil {
		return err
	}
	return serverStream.Send(response)
}

func (h *_HelloYARPCHandler) HelloInStream(serverStream *protobuf.ServerStream) error {
	requestMessage, err := serverStream.Receive(newHelloServiceHelloInStreamYARPCRequest)
	if requestMessage == nil {
		return err
	}

	request, ok := requestMessage.(*HelloRequest)
	if !ok {
		return protobuf.CastError(emptyHelloServiceHelloInStreamYARPCRequest, requestMessage)
	}
	return h.server.HelloInStream(request, &_HelloServiceHelloInStreamYARPCServer{serverStream: serverStream})
}

func (h *_HelloYARPCHandler) HelloThere(serverStream *protobuf.ServerStream) error {
	return h.server.HelloThere(&_HelloServiceHelloThereYARPCServer{serverStream: serverStream})
}

type _HelloServiceHelloOutStreamYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_HelloServiceHelloOutStreamYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_HelloServiceHelloOutStreamYARPCClient) Send(request *HelloRequest, options ...yarpc.StreamOption) error {
	return c.stream.Send(request, options...)
}

func (c *_HelloServiceHelloOutStreamYARPCClient) CloseAndRecv(options ...yarpc.StreamOption) (*HelloResponse, error) {
	if err := c.stream.Close(options...); err != nil {
		return nil, err
	}
	responseMessage, err := c.stream.Receive(newHelloServiceHelloOutStreamYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HelloResponse)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloOutStreamYARPCResponse, responseMessage)
	}
	return response, err
}

type _HelloServiceHelloInStreamYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_HelloServiceHelloInStreamYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_HelloServiceHelloInStreamYARPCClient) Recv(options ...yarpc.StreamOption) (*HelloResponse, error) {
	responseMessage, err := c.stream.Receive(newHelloServiceHelloInStreamYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HelloResponse)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloInStreamYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_HelloServiceHelloInStreamYARPCClient) CloseSend(options ...yarpc.StreamOption) error {
	return c.stream.Close(options...)
}

type _HelloServiceHelloThereYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_HelloServiceHelloThereYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_HelloServiceHelloThereYARPCClient) Send(request *HelloRequest, options ...yarpc.StreamOption) error {
	return c.stream.Send(request, options...)
}

func (c *_HelloServiceHelloThereYARPCClient) Recv(options ...yarpc.StreamOption) (*HelloResponse, error) {
	responseMessage, err := c.stream.Receive(newHelloServiceHelloThereYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HelloResponse)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloThereYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_HelloServiceHelloThereYARPCClient) CloseSend(options ...yarpc.StreamOption) error {
	return c.stream.Close(options...)
}

type _HelloServiceHelloOutStreamYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_HelloServiceHelloOutStreamYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_HelloServiceHelloOutStreamYARPCServer) Recv(options ...yarpc.StreamOption) (*HelloRequest, error) {
	requestMessage, err := s.serverStream.Receive(newHelloServiceHelloOutStreamYARPCRequest, options...)
	if requestMessage == nil {
		return nil, err
	}
	request, ok := requestMessage.(*HelloRequest)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloOutStreamYARPCRequest, requestMessage)
	}
	return request, err
}

type _HelloServiceHelloInStreamYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_HelloServiceHelloInStreamYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_HelloServiceHelloInStreamYARPCServer) Send(response *HelloResponse, options ...yarpc.StreamOption) error {
	return s.serverStream.Send(response, options...)
}

type _HelloServiceHelloThereYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_HelloServiceHelloThereYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_HelloServiceHelloThereYARPCServer) Recv(options ...yarpc.StreamOption) (*HelloRequest, error) {
	requestMessage, err := s.serverStream.Receive(newHelloServiceHelloThereYARPCRequest, options...)
	if requestMessage == nil {
		return nil, err
	}
	request, ok := requestMessage.(*HelloRequest)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloThereYARPCRequest, requestMessage)
	}
	return request, err
}

func (s *_HelloServiceHelloThereYARPCServer) Send(response *HelloResponse, options ...yarpc.StreamOption) error {
	return s.serverStream.Send(response, options...)
}

func newHelloServiceHelloUnaryYARPCRequest() proto.Message {
	return &HelloRequest{}
}

func newHelloServiceHelloUnaryYARPCResponse() proto.Message {
	return &HelloResponse{}
}

func newHelloServiceHelloThereYARPCRequest() proto.Message {
	return &HelloRequest{}
}

func newHelloServiceHelloThereYARPCResponse() proto.Message {
	return &HelloResponse{}
}

func newHelloServiceHelloOutStreamYARPCRequest() proto.Message {
	return &HelloRequest{}
}

func newHelloServiceHelloOutStreamYARPCResponse() proto.Message {
	return &HelloResponse{}
}

func newHelloServiceHelloInStreamYARPCRequest() proto.Message {
	return &HelloRequest{}
}

func newHelloServiceHelloInStreamYARPCResponse() proto.Message {
	return &HelloResponse{}
}

var (
	emptyHelloServiceHelloUnaryYARPCRequest      = &HelloRequest{}
	emptyHelloServiceHelloUnaryYARPCResponse     = &HelloResponse{}
	emptyHelloServiceHelloThereYARPCRequest      = &HelloRequest{}
	emptyHelloServiceHelloThereYARPCResponse     = &HelloResponse{}
	emptyHelloServiceHelloOutStreamYARPCRequest  = &HelloRequest{}
	emptyHelloServiceHelloOutStreamYARPCResponse = &HelloResponse{}
	emptyHelloServiceHelloInStreamYARPCRequest   = &HelloRequest{}
	emptyHelloServiceHelloInStreamYARPCResponse  = &HelloResponse{}
)

var yarpcFileDescriptorClosure45d12c3ddf34baf8 = [][]byte{
	// internal/examples/streaming/stream.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0xcc, 0x2b, 0x49,
		0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x2e,
		0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0x87, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
		0xd4, 0x4a, 0x93, 0x52, 0x8b, 0xf4, 0x2a, 0x13, 0x8b, 0x0a, 0x92, 0xf5, 0x60, 0x9a, 0xf4, 0x60,
		0x9a, 0xf4, 0xe0, 0x9a, 0x94, 0xe4, 0xb8, 0x78, 0x3c, 0x52, 0x73, 0x72, 0xf2, 0x83, 0x52, 0x0b,
		0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
		0x83, 0x98, 0x32, 0x53, 0x94, 0xe4, 0xb9, 0x78, 0xa1, 0xf2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
		0xe8, 0x0a, 0x8c, 0x7a, 0x58, 0xb8, 0x58, 0xc1, 0x2a, 0x84, 0xaa, 0xb9, 0xb8, 0xc0, 0x8c, 0xd0,
		0xbc, 0xc4, 0xa2, 0x4a, 0x21, 0x13, 0x3d, 0xe2, 0x5c, 0xa0, 0x87, 0x6c, 0xbd, 0x94, 0x29, 0x89,
		0xba, 0x20, 0x8e, 0x52, 0x62, 0x10, 0xaa, 0x87, 0x5a, 0x1e, 0x92, 0x91, 0x5a, 0x94, 0x4a, 0x67,
		0xcb, 0x35, 0x18, 0x0d, 0x18, 0x85, 0x1a, 0x19, 0xb9, 0xf8, 0xc0, 0xe2, 0xfe, 0xa5, 0x25, 0xc1,
		0x60, 0x95, 0x74, 0x77, 0x85, 0x50, 0x03, 0x23, 0x34, 0xb6, 0x3c, 0xf3, 0x06, 0xc4, 0x09, 0x06,
		0x8c, 0x4e, 0xdc, 0x51, 0x9c, 0x70, 0xe9, 0x24, 0x36, 0x70, 0x5a, 0x34, 0x06, 0x04, 0x00, 0x00,
		0xff, 0xff, 0xef, 0x5c, 0x66, 0xbe, 0xb7, 0x02, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) HelloYARPCClient {
			return NewHelloYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}

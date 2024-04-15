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

package secondserviceserver

import (
	context "context"
	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	gauntlet "go.uber.org/yarpc/internal/crossdock/thrift/gauntlet"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the SecondService service.
type Interface interface {
	BlahBlah(
		ctx context.Context,
	) error

	SecondtestString(
		ctx context.Context,
		Thing *string,
	) (string, error)
}

// New prepares an implementation of the SecondService service for
// registration.
//
//	handler := SecondServiceHandler{}
//	dispatcher.Register(secondserviceserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "SecondService",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "blahBlah",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Unary,
					Unary:  thrift.UnaryHandler(h.BlahBlah),
					NoWire: blahblah_NoWireHandler{impl},
				},
				Signature:    "BlahBlah()",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "secondtestString",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Unary,
					Unary:  thrift.UnaryHandler(h.SecondtestString),
					NoWire: secondteststring_NoWireHandler{impl},
				},
				Signature:    "SecondtestString(Thing *string) (string)",
				ThriftModule: gauntlet.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 2)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type yarpcErrorNamer interface{ YARPCErrorName() string }

type yarpcErrorCoder interface{ YARPCErrorCode() *yarpcerrors.Code }

func (h handler) BlahBlah(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.SecondService_BlahBlah_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'SecondService' procedure 'BlahBlah': %w", err)
	}

	appErr := h.impl.BlahBlah(ctx)

	hadError := appErr != nil
	result, err := gauntlet.SecondService_BlahBlah_Helper.WrapResponse(appErr)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}

	return response, err
}

func (h handler) SecondtestString(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.SecondService_SecondtestString_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'SecondService' procedure 'SecondtestString': %w", err)
	}

	success, appErr := h.impl.SecondtestString(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.SecondService_SecondtestString_Helper.WrapResponse(success, appErr)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}

	return response, err
}

type blahblah_NoWireHandler struct{ impl Interface }

func (h blahblah_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args gauntlet.SecondService_BlahBlah_Args
		rw   stream.ResponseWriter
		err  error
	)

	rw, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args)
	if err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'SecondService' procedure 'BlahBlah': %w", err)
	}

	appErr := h.impl.BlahBlah(ctx)

	hadError := appErr != nil
	result, err := gauntlet.SecondService_BlahBlah_Helper.WrapResponse(appErr)
	response := thrift.NoWireResponse{ResponseWriter: rw}
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}
	return response, err

}

type secondteststring_NoWireHandler struct{ impl Interface }

func (h secondteststring_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args gauntlet.SecondService_SecondtestString_Args
		rw   stream.ResponseWriter
		err  error
	)

	rw, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args)
	if err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'SecondService' procedure 'SecondtestString': %w", err)
	}

	success, appErr := h.impl.SecondtestString(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.SecondService_SecondtestString_Helper.WrapResponse(success, appErr)
	response := thrift.NoWireResponse{ResponseWriter: rw}
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}
	return response, err

}

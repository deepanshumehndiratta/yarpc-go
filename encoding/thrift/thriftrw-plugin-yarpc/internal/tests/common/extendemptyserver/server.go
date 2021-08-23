// Code generated by thriftrw-plugin-yarpc
// @generated

package extendemptyserver

import (
	context "context"
	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	common "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common"
	emptyserviceserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common/emptyserviceserver"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the ExtendEmpty service.
type Interface interface {
	emptyserviceserver.Interface

	Hello(
		ctx context.Context,
	) error
}

// New prepares an implementation of the ExtendEmpty service for
// registration.
//
// 	handler := ExtendEmptyHandler{}
// 	dispatcher.Register(extendemptyserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "ExtendEmpty",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "hello",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Unary,
					Unary:  thrift.UnaryHandler(h.Hello),
					NoWire: hello_NoWireHandler{impl},
				},
				Signature:    "Hello()",
				ThriftModule: common.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)

	procedures = append(
		procedures,
		emptyserviceserver.New(
			impl,
			append(
				opts,
				thrift.Named("ExtendEmpty"),
			)...,
		)...,
	)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type yarpcErrorNamer interface{ YARPCErrorName() string }

type yarpcErrorCoder interface{ YARPCErrorCode() *yarpcerrors.Code }

func (h handler) Hello(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args common.ExtendEmpty_Hello_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ExtendEmpty' procedure 'Hello': %w", err)
	}

	appErr := h.impl.Hello(ctx)

	hadError := appErr != nil
	result, err := common.ExtendEmpty_Hello_Helper.WrapResponse(appErr)

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

type hello_NoWireHandler struct{ impl Interface }

func (h hello_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args common.ExtendEmpty_Hello_Args
		rw   stream.ResponseWriter
		err  error
	)

	rw, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args)
	if err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'ExtendEmpty' procedure 'Hello': %w", err)
	}

	appErr := h.impl.Hello(ctx)

	hadError := appErr != nil
	result, err := common.ExtendEmpty_Hello_Helper.WrapResponse(appErr)
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

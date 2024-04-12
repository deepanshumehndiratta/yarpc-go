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

package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"go.uber.org/thriftrw/compile"
	"go.uber.org/thriftrw/plugin"
)

const (
	_errorCodeAnnotationKey = "rpc.code"
)

const yarpcerrorTemplate = `
// Code generated by thriftrw-plugin-yarpc
// @generated

package <.Name>

<range $key, $val := .Types>
	<if (isException $val)>
	<$yarpcerrors := import "go.uber.org/yarpc/yarpcerrors" ->
	// YARPCErrorCode returns <if isSetYARPCCode .Annotations>a <getYARPCErrorCode .><else>nil<end> for <$val.Name>.
	//
	// This is derived from the rpc.code annotation on the Thrift exception.
	func (e *<$val.Name>) YARPCErrorCode() *<$yarpcerrors>.Code {
		<if isSetYARPCCode .Annotations>code := <getYARPCErrorCode .>
		return &code
		<else>
		return nil
		<end>}

	// Name is the error name for <$val.Name>.
	func (e *<$val.Name>) YARPCErrorName() string { return <getYARPCErrorName .> }
	<end>
<end>
`

var (
	_gRPCCodeNameToYARPCErrorCodeType = map[string]string{
		// https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
		"CANCELLED":           "yarpcerrors.CodeCancelled",
		"UNKNOWN":             "yarpcerrors.CodeUnknown",
		"INVALID_ARGUMENT":    "yarpcerrors.CodeInvalidArgument",
		"DEADLINE_EXCEEDED":   "yarpcerrors.CodeDeadlineExceeded",
		"NOT_FOUND":           "yarpcerrors.CodeNotFound",
		"ALREADY_EXISTS":      "yarpcerrors.CodeAlreadyExists",
		"PERMISSION_DENIED":   "yarpcerrors.CodePermissionDenied",
		"RESOURCE_EXHAUSTED":  "yarpcerrors.CodeResourceExhausted",
		"FAILED_PRECONDITION": "yarpcerrors.CodeFailedPrecondition",
		"ABORTED":             "yarpcerrors.CodeAborted",
		"OUT_OF_RANGE":        "yarpcerrors.CodeOutOfRange",
		"UNIMPLEMENTED":       "yarpcerrors.CodeUnimplemented",
		"INTERNAL":            "yarpcerrors.CodeInternal",
		"UNAVAILABLE":         "yarpcerrors.CodeUnavailable",
		"DATA_LOSS":           "yarpcerrors.CodeDataLoss",
		"UNAUTHENTICATED":     "yarpcerrors.CodeUnauthenticated",
	}

	_availableCodes = fmt.Sprintf(`Available codes: %s`, strings.Join(
		// Codes are listed below in enum-order, derived from:
		// - https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
		[]string{
			"CANCELLED",
			"UNKNOWN",
			"INVALID_ARGUMENT",
			"DEADLINE_EXCEEDED",
			"NOT_FOUND",
			"ALREADY_EXISTS",
			"PERMISSION_DENIED",
			"RESOURCE_EXHAUSTED",
			"FAILED_PRECONDITION",
			"ABORTED",
			"OUT_OF_RANGE",
			"UNIMPLEMENTED",
			"INTERNAL",
			"UNAVAILABLE",
			"DATA_LOSS",
			"UNAUTHENTICATED",
		}, ","))
)

func yarpcErrorGenerator(data *moduleTemplateData, files map[string][]byte) error {
	// kv.thrift => .../kv/types_yarpc.go
	path := filepath.Join(data.Module.Directory, "types_yarpc.go")

	// Get the original thrift file path
	thriftFilePath := data.Module.GetThriftFilePath()
	// and re-compile it. Plugins do not actually have access to the original
	// thrift file's module, and this is how we gain access to it, since this
	// generator writes extra methods to exception types defined in the Thrift
	// file.
	compiledModule, err := compile.Compile(thriftFilePath)
	if err != nil {
		return fmt.Errorf("error compiling the thrift file: %s", err.Error())
	}

	templateOptions := append(templateOptions,
		plugin.TemplateFunc("isException", func(t compile.TypeSpec) bool {
			if t, ok := t.(*compile.StructSpec); ok {
				return t.IsExceptionType()
			}
			return false
		}),
		plugin.TemplateFunc("isSetYARPCCode", func(a compile.Annotations) bool {
			_, found := a[_errorCodeAnnotationKey]
			return found
		}),
		plugin.TemplateFunc("getYARPCErrorCode", getYARPCErrorCode),
		plugin.TemplateFunc("getYARPCErrorName", getYARPCErrorName),
	)

	files[path], err = plugin.GoFileFromTemplate(path, yarpcerrorTemplate, compiledModule, templateOptions...)
	return err
}

func getYARPCErrorCode(t *compile.StructSpec) string {
	errorCodeString := t.Annotations[_errorCodeAnnotationKey]
	yCode, ok := _gRPCCodeNameToYARPCErrorCodeType[errorCodeString]
	if !ok {
		panic(fmt.Sprintf("invalid rpc.code annotation for %q: %q\n%s", t.Name, errorCodeString, _availableCodes))
	}
	return yCode
}

func getYARPCErrorName(t *compile.StructSpec) string {
	return fmt.Sprintf("%q", t.ThriftName())
}

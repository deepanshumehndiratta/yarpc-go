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

package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"
	"go.uber.org/net/metrics"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/internal/observability"
	"go.uber.org/yarpc/yarpcerrors"
	"google.golang.org/grpc/metadata"
)

func TestResponseWriterAddHeaders(t *testing.T) {
	tests := map[string]struct {
		h               transport.Headers
		md              metadata.MD
		expErr          error
		expReportHeader bool
		expMD           metadata.MD
	}{
		"md-is-nil": {
			h:     transport.NewHeaders().With("foo", "bar"),
			md:    nil,
			expMD: metadata.Pairs("foo", "bar"),
		},
		"success": {
			h:     transport.NewHeaders().With("foo", "bar"),
			md:    metadata.Pairs(),
			expMD: metadata.Pairs("foo", "bar"),
		},
		"reserved-header-used": {
			h:      transport.NewHeaders().With("rpc-any", "any-value"),
			md:     metadata.Pairs(),
			expErr: yarpcerrors.InvalidArgumentErrorf("cannot use reserved header in application headers: rpc-any"),
			expMD:  metadata.Pairs(),
		},
		"report-header": {
			h:               transport.NewHeaders().With("$rpc$-any", "any-value"),
			md:              metadata.Pairs(),
			expMD:           metadata.Pairs("$rpc$-any", "any-value"),
			expReportHeader: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			root := metrics.New()
			m := observability.NewReserveHeaderMetrics(root.Scope(), "test")
			rw := newResponseWriter(m.With("any-source", "any-test"))
			rw.md = tt.md

			rw.AddHeaders(tt.h)
			if tt.expErr != nil {
				errs := multierr.Errors(rw.headerErr)
				require.Len(t, errs, 1)
				assert.Equal(t, tt.expErr, errs[0])
			} else {
				assert.NoError(t, rw.headerErr)
			}
			assert.Equal(t, tt.expMD, rw.md)

			if tt.expReportHeader {
				assertTuple(t, root.Snapshot().Counters, tuple{"test_reserved_headers_error", "any-source", "any-test", 1})
			} else {
				assertEmptyMetrics(t, root.Snapshot())
			}
		})
	}
}

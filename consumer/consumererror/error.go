// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumererror // import "go.opentelemetry.io/collector/consumer/consumererror"

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.opentelemetry.io/collector/consumer/consumererror/internal/statusconversion"
)

// Error is intended to be used to encapsulate various information that can add
// context to an error that occurred within a pipeline component. Error objects
// are constructed through calling `New` with the relevant options to capture
// data around the error that occurred.
//
// It may hold multiple errors from downstream components, and can be merged
// with other errors as it travels upstream using `Combine`. The `Error` should
// be obtained from a given `error` object using `errors.As`.
//
// Experimental: This API is at the early stage of development and may change
// without backward compatibility
type Error struct {
	error
	httpStatus int
	grpcStatus *status.Status
	retryable  bool
}

var _ error = (*Error)(nil)

// ErrorOption allows annotating an Error with metadata.
type ErrorOption interface {
	applyOption(*Error)
}

type errorOptionFunc func(*Error)

func (f errorOptionFunc) applyOption(e *Error) {
	f(e)
}

// New wraps an error that happened while consuming telemetry and adds metadata
// onto it to be passed back up the pipeline.
// At least one option should be provided.
//
// Experimental: This API is at the early stage of development and may change
// without backward compatibility
func New(origErr error, options ...ErrorOption) error {
	err := &Error{error: origErr}

	for _, option := range options {
		option.applyOption(err)
	}

	return err
}

// WithOTLPHTTPStatus records an HTTP status code that was received from a server
// during data submission.
// It is not necessary to use WithRetryable with creating an error with WithOTLPHTTPStatus
// as the retryable property can be inferred from the HTTP status code using OTLP specification.
//
// Experimental: This API is at the early stage of development and may change
// without backward compatibility
func WithOTLPHTTPStatus(status int) ErrorOption {
	return errorOptionFunc(func(err *Error) {
		err.httpStatus = status
	})
}

// WithOTLPGRPCStatus records a gRPC status code that was received from a server
// during data submission.
// It is not necessary to use WithRetryable with creating an error with WithOTLPGRPCStatus
// as the retryable property can be inferred from the grpc status using OTLP specification.
//
// Experimental: This API is at the early stage of development and may change
// without backward compatibility
func WithOTLPGRPCStatus(status *status.Status) ErrorOption {
	return errorOptionFunc(func(err *Error) {
		err.grpcStatus = status
	})
}

// WithRetryable records that this error is retryable according to OTLP specification.
// WithRetryable is not necessary when creating an error with WithOTLPHTTPStatus or
// WithOTLPGRPCStatus, as the retryable property can be inferred from OTLP specification.
//
// Experimental: This API is at the early stage of development and may change
// without backward compatibility
func WithRetryable() ErrorOption {
	return errorOptionFunc(func(err *Error) {
		err.retryable = true
	})
}

// Error implements the error interface.
func (e *Error) Error() string {
	return e.error.Error()
}

// Unwrap returns the wrapped error for use by `errors.Is` and `errors.As`.
func (e *Error) Unwrap() error {
	return e.error
}

// OTLPHTTPStatus returns an HTTP status code either directly set by the source,
// derived from a gRPC status code set by the source, or derived from Retryable.
// When deriving the value, the OTLP specification is used to map to HTTP.
// See https://github.com/open-telemetry/opentelemetry-proto/blob/main/docs/specification.md for more details.
//
// If a http status code cannot be derived from these three sources then 500 is returned.
//
// Experimental: This API is at the early stage of development and may change
// without backward compatibility
func (e *Error) OTLPHTTPStatus() int {
	if e.httpStatus != 0 {
		return e.httpStatus
	}
	if e.grpcStatus != nil {
		return statusconversion.GetHTTPStatusCodeFromStatus(e.grpcStatus)
	}
	if e.retryable {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

// OTLPGRPCStatus returns an gRPC status code either directly set by the source,
// derived from an HTTP status code set by the source, or derived from Retryable.
// When deriving the value, the OTLP specification is used to map to GRPC.
// See https://github.com/open-telemetry/opentelemetry-proto/blob/main/docs/specification.md for more details.
//
// If a grpc code cannot be derived from these three sources then INTERNAL is returned.
//
// Experimental: This API is at the early stage of development and may change
// without backward compatibility
func (e *Error) OTLPGRPCStatus() *status.Status {
	if e.grpcStatus != nil {
		return e.grpcStatus
	}
	if e.httpStatus != 0 {
		return statusconversion.NewStatusFromMsgAndHTTPCode(e.Error(), e.httpStatus)
	}
	if e.retryable {
		return status.New(codes.Unavailable, e.Error())
	}
	return status.New(codes.Internal, e.Error())
}

// Retryable returns true if the error was created with the WithRetryable set to true,
// if the http status code is retryable according to OTLP,
// or if the grpc status is retryable according to OTLP.
// Otherwise, returns false.
//
// See https://github.com/open-telemetry/opentelemetry-proto/blob/main/docs/specification.md for retryable
// http and grpc codes.
//
// Experimental: This API is at the early stage of development and may change
// without backward compatibility
func (e *Error) Retryable() bool {
	if e.retryable {
		return true
	}
	switch e.httpStatus {
	case http.StatusTooManyRequests, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return true
	}
	if e.grpcStatus != nil {
		switch e.grpcStatus.Code() {
		case codes.Canceled, codes.DeadlineExceeded, codes.Aborted, codes.OutOfRange, codes.Unavailable, codes.DataLoss:
			return true
		}
	}
	return false
}

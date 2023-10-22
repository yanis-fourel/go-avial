package hgtp

import (
	"fmt"
)

type AvesTerraError struct {
	Code    Error
	Message string
}

func (e *AvesTerraError) Error() string {
	// TODO: print error code name as string instead of number
	return fmt.Sprintf("AvesTerra error %d: %s", e.Code, e.Message)
}

// Default adapter error
func NewErrorApplication(msg string) *AvesTerraError {
	return &AvesTerraError{Code: ErrorApplication, Message: msg}
}

// Returned by the adapter callback.
// The `ErrorBypass` error indicates the adapter did not process the
// invocation request, and is deferring the operation to the next lower
// precedence level adapter, if any, so that it may also process the
// request
// see `ErrorBypass` for more information
func NewErrorBypass(msg string) *AvesTerraError {
	return &AvesTerraError{Code: ErrorBypass, Message: msg}
}

// Returned by the adapter callback.
// The `ErrorForward` error indicates that the adapter did successfully
// fulfil the invocation request, but that it is forwarding it on to the
// next lower precedence level adapter, if any, so that it may also process
// the request.
// see `ErrorForward` for more information
func NewErrorForward(msg string) *AvesTerraError {
	return &AvesTerraError{Code: ErrorForward, Message: msg}
}

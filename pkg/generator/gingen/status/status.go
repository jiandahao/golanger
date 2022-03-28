package status

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

// Reference: google.golang.org/grpc/codes/codes.go
var (
	// OK is returned on success.
	OK = NewErrorCode(codes.OK, http.StatusOK, "ok")

	// Canceled indicates the operation was canceled (typically by the caller).
	Canceled = NewErrorCode(codes.Canceled, http.StatusRequestTimeout, "cancled")

	// Unknown error. An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Unknown = NewErrorCode(codes.Unknown, http.StatusInternalServerError, "unknown")

	// InvalidArgument indicates client specified an invalid argument.
	// Note that this differs from FailedPrecondition. It indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	InvalidArgument = NewErrorCode(codes.InvalidArgument, http.StatusBadRequest, "invalid_argument")

	// DeadlineExceeded means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	DeadlineExceeded = NewErrorCode(codes.DeadlineExceeded, http.StatusGatewayTimeout, "deadline_exceeded")

	// NotFound means some requested entity (e.g., file or directory) was
	// not found.
	NotFound = NewErrorCode(codes.NotFound, http.StatusNotFound, "not_found")

	// AlreadyExists means an attempt to create an entity failed because one
	// already exists.
	AlreadyExists = NewErrorCode(codes.AlreadyExists, http.StatusConflict, "already_exists")

	// PermissionDenied indicates the caller does not have permission to
	// execute the specified operation. It must not be used for rejections
	// caused by exhausting some resource (use ResourceExhausted
	// instead for those errors). It must not be
	// used if the caller cannot be identified (use Unauthenticated
	// instead for those errors).
	PermissionDenied = NewErrorCode(codes.PermissionDenied, http.StatusForbidden, "permission_denied")

	// ResourceExhausted indicates some resource has been exhausted, perhaps
	// a per-user quota, or perhaps the entire file system is out of space.
	ResourceExhausted = NewErrorCode(codes.ResourceExhausted, http.StatusTooManyRequests, "resource_exhausted")

	// FailedPrecondition indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	// For example, directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FailedPrecondition, Aborted, and Unavailable:
	//  (a) Use Unavailable if the client can retry just the failing call.
	//  (b) Use Aborted if the client should retry at a higher-level
	//      (e.g., restarting a read-modify-write sequence).
	//  (c) Use FailedPrecondition if the client should not retry until
	//      the system state has been explicitly fixed. E.g., if an "rmdir"
	//      fails because the directory is non-empty, FailedPrecondition
	//      should be returned since the client should not retry unless
	//      they have first fixed up the directory by deleting files from it.
	//  (d) Use FailedPrecondition if the client performs conditional
	//      REST Get/Update/Delete on a resource and the resource on the
	//      server does not match the condition. E.g., conflicting
	//      read-modify-write on the same resource.
	FailedPrecondition = NewErrorCode(codes.FailedPrecondition, http.StatusBadRequest, "failed_precondition")

	// Aborted indicates the operation was aborted, typically due to a
	// concurrency issue like sequencer check failures, transaction aborts,
	// etc.
	//
	// See litmus test above for deciding between FailedPrecondition,
	// Aborted, and Unavailable.
	Aborted = NewErrorCode(codes.Aborted, http.StatusConflict, "aborted")

	// OutOfRange means operation was attempted past the valid range.
	// E.g., seeking or reading past end of file.
	//
	// Unlike InvalidArgument, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate InvalidArgument if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// OutOfRange if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between FailedPrecondition and
	// OutOfRange. We recommend using OutOfRange (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an OutOfRange error to detect when
	// they are done.
	OutOfRange = NewErrorCode(codes.OutOfRange, http.StatusBadRequest, "out_of_range")

	// Unimplemented indicates operation is not implemented or not
	// supported/enabled in this service.
	Unimplemented = NewErrorCode(codes.Unimplemented, http.StatusNotImplemented, "unimplement")

	// Internal errors. Means some invariants expected by underlying
	// system has been broken. If you see one of these errors,
	// something is very broken.
	Internal = NewErrorCode(codes.Internal, http.StatusInternalServerError, "internal")

	// Unavailable indicates the service is currently unavailable.
	// This is a most likely a transient condition and may be corrected
	// by retrying with a backoff. Note that it is not always safe to retry
	// non-idempotent operations.
	Unavailable = NewErrorCode(codes.Unavailable, http.StatusServiceUnavailable, "unavailable")

	// DataLoss indicates unrecoverable data loss or corruption.
	DataLoss = NewErrorCode(codes.DataLoss, http.StatusInternalServerError, "data_loss")

	// Unauthenticated indicates the request does not have valid
	// authentication credentials for the operation.
	Unauthenticated = NewErrorCode(codes.Unauthenticated, http.StatusUnauthorized, "unauthenticated")
)

// ErrorDetails error with status
type ErrorDetails struct {
	Code    codes.Code `json:"code,omitempty"`
	Status  int        `json:"status,omitempty"` // http status code
	Msg     string     `json:"msg,omitempty"`
	Details string     `json:"details,omitempty"`
}

func (e *ErrorDetails) Error() string {
	return fmt.Sprintf("error: code = %v, status = %v, msg = %s, details = %s", e.Code, e.Status, e.Msg, e.Details)
}

var errorMap = make(map[codes.Code]*ErrorDetails)

// NewErrorCode creates and registers an error code
func NewErrorCode(code codes.Code, status int /*http status code*/, msg string) codes.Code {
	if _, ok := errorMap[code]; ok {
		panic("runtime: duplicated error code")
	}

	errorMap[code] = &ErrorDetails{Code: code, Status: status, Msg: msg}
	return code
}

// FromCode find registered error by error code.
func FromCode(code codes.Code) (*ErrorDetails, bool) {
	if e, ok := errorMap[code]; ok {
		return e, true
	}

	return nil, false
}

// Error returns an error representing c and msg.  If c is OK, returns nil.
func Error(c codes.Code, msg string) error {
	if c == OK {
		return nil
	}

	e, ok := errorMap[c]
	if !ok {
		return &ErrorDetails{Code: c, Msg: "unrecongnized_error_code", Details: msg}
	}
	return &ErrorDetails{Code: c, Msg: e.Msg, Status: e.Status, Details: msg}
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c codes.Code, format string, a ...interface{}) error {
	return Error(c, fmt.Sprintf(format, a...))
}

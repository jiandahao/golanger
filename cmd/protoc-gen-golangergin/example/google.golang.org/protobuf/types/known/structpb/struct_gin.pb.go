// Code generated. DO NOT EDIT.

package structpb

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	runtime "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// `Struct` represents a structured data value, consisting of fields
// which map to dynamically typed values. In some languages, `Struct`
// might be supported by a native representation. For example, in
// scripting languages like JS a struct is represented as an
// object. The details of that representation are described together
// with the proto support for the language.
//
// The JSON representation for `Struct` is JSON object.
type Struct struct {
	Fields map[string]*Value `json:"fields,omitempty"`
}

// `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of these
// variants. Absence of any variant indicates an error.
//
// The JSON representation for `Value` is JSON value.
type Value struct {
	NullValue   *NullValue `json:"null_value,omitempty"`
	NumberValue *float64   `json:"number_value,omitempty"`
	StringValue *string    `json:"string_value,omitempty"`
	BoolValue   *bool      `json:"bool_value,omitempty"`
	StructValue *Struct    `json:"struct_value,omitempty"`
	ListValue   *ListValue `json:"list_value,omitempty"`
}

// `ListValue` is a wrapper around a repeated field of values.
//
// The JSON representation for `ListValue` is JSON array.
type ListValue struct {
	Values []*Value `json:"values,omitempty"`
}

type Struct_FieldsEntry struct {
	Key   string `json:"key,omitempty"`
	Value *Value `json:"value,omitempty"`
}

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//  The JSON representation for `NullValue` is JSON `null`.
type NullValue int32

const (
	// Null value.
	NullValue_NULL_VALUE NullValue = 0
)

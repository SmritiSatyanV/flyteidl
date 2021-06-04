// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/core/types.proto

package core

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _types_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SchemaType with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SchemaType) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetColumns() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SchemaTypeValidationError{
					field:  fmt.Sprintf("Columns[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SchemaTypeValidationError is the validation error returned by
// SchemaType.Validate if the designated constraints aren't met.
type SchemaTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaTypeValidationError) ErrorName() string { return "SchemaTypeValidationError" }

// Error satisfies the builtin error interface
func (e SchemaTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchemaType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaTypeValidationError{}

// Validate checks the field values on BlobType with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *BlobType) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Format

	// no validation rules for Dimensionality

	return nil
}

// BlobTypeValidationError is the validation error returned by
// BlobType.Validate if the designated constraints aren't met.
type BlobTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlobTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlobTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlobTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlobTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlobTypeValidationError) ErrorName() string { return "BlobTypeValidationError" }

// Error satisfies the builtin error interface
func (e BlobTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlobType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlobTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlobTypeValidationError{}

// Validate checks the field values on EnumType with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *EnumType) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DefaultValue

	return nil
}

// EnumTypeValidationError is the validation error returned by
// EnumType.Validate if the designated constraints aren't met.
type EnumTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnumTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnumTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnumTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnumTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnumTypeValidationError) ErrorName() string { return "EnumTypeValidationError" }

// Error satisfies the builtin error interface
func (e EnumTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnumType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnumTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnumTypeValidationError{}

// Validate checks the field values on LiteralType with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LiteralType) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LiteralTypeValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Type.(type) {

	case *LiteralType_Simple:
		// no validation rules for Simple

	case *LiteralType_Schema:

		if v, ok := interface{}(m.GetSchema()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "Schema",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralType_CollectionType:

		if v, ok := interface{}(m.GetCollectionType()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "CollectionType",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralType_MapValueType:

		if v, ok := interface{}(m.GetMapValueType()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "MapValueType",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralType_Blob:

		if v, ok := interface{}(m.GetBlob()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "Blob",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralType_EnumType:

		if v, ok := interface{}(m.GetEnumType()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralTypeValidationError{
					field:  "EnumType",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LiteralTypeValidationError is the validation error returned by
// LiteralType.Validate if the designated constraints aren't met.
type LiteralTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LiteralTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LiteralTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LiteralTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LiteralTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LiteralTypeValidationError) ErrorName() string { return "LiteralTypeValidationError" }

// Error satisfies the builtin error interface
func (e LiteralTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLiteralType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LiteralTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LiteralTypeValidationError{}

// Validate checks the field values on OutputReference with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OutputReference) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NodeId

	// no validation rules for Var

	return nil
}

// OutputReferenceValidationError is the validation error returned by
// OutputReference.Validate if the designated constraints aren't met.
type OutputReferenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutputReferenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutputReferenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutputReferenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutputReferenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutputReferenceValidationError) ErrorName() string { return "OutputReferenceValidationError" }

// Error satisfies the builtin error interface
func (e OutputReferenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutputReference.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutputReferenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutputReferenceValidationError{}

// Validate checks the field values on Error with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Error) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FailedNodeId

	// no validation rules for Message

	return nil
}

// ErrorValidationError is the validation error returned by Error.Validate if
// the designated constraints aren't met.
type ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorValidationError) ErrorName() string { return "ErrorValidationError" }

// Error satisfies the builtin error interface
func (e ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorValidationError{}

// Validate checks the field values on SchemaType_SchemaColumn with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SchemaType_SchemaColumn) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Type

	return nil
}

// SchemaType_SchemaColumnValidationError is the validation error returned by
// SchemaType_SchemaColumn.Validate if the designated constraints aren't met.
type SchemaType_SchemaColumnValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaType_SchemaColumnValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaType_SchemaColumnValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaType_SchemaColumnValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaType_SchemaColumnValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaType_SchemaColumnValidationError) ErrorName() string {
	return "SchemaType_SchemaColumnValidationError"
}

// Error satisfies the builtin error interface
func (e SchemaType_SchemaColumnValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchemaType_SchemaColumn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaType_SchemaColumnValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaType_SchemaColumnValidationError{}

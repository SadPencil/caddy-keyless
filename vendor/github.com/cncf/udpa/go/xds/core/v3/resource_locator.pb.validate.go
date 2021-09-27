// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: xds/core/v3/resource_locator.proto

package xds_core_v3

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
var _resource_locator_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ResourceLocator with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ResourceLocator) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := ResourceLocator_Scheme_name[int32(m.GetScheme())]; !ok {
		return ResourceLocatorValidationError{
			field:  "Scheme",
			reason: "value must be one of the defined enum values",
		}
	}

	// no validation rules for Id

	// no validation rules for Authority

	if utf8.RuneCountInString(m.GetResourceType()) < 1 {
		return ResourceLocatorValidationError{
			field:  "ResourceType",
			reason: "value length must be at least 1 runes",
		}
	}

	for idx, item := range m.GetDirectives() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceLocatorValidationError{
					field:  fmt.Sprintf("Directives[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.ContextParamSpecifier.(type) {

	case *ResourceLocator_ExactContext:

		if v, ok := interface{}(m.GetExactContext()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceLocatorValidationError{
					field:  "ExactContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ResourceLocatorValidationError is the validation error returned by
// ResourceLocator.Validate if the designated constraints aren't met.
type ResourceLocatorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceLocatorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceLocatorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceLocatorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceLocatorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceLocatorValidationError) ErrorName() string { return "ResourceLocatorValidationError" }

// Error satisfies the builtin error interface
func (e ResourceLocatorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceLocator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceLocatorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceLocatorValidationError{}

// Validate checks the field values on ResourceLocator_Directive with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResourceLocator_Directive) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Directive.(type) {

	case *ResourceLocator_Directive_Alt:

		if v, ok := interface{}(m.GetAlt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceLocator_DirectiveValidationError{
					field:  "Alt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ResourceLocator_Directive_Entry:

		if utf8.RuneCountInString(m.GetEntry()) < 1 {
			return ResourceLocator_DirectiveValidationError{
				field:  "Entry",
				reason: "value length must be at least 1 runes",
			}
		}

		if !_ResourceLocator_Directive_Entry_Pattern.MatchString(m.GetEntry()) {
			return ResourceLocator_DirectiveValidationError{
				field:  "Entry",
				reason: "value does not match regex pattern \"^[0-9a-zA-Z_\\\\-\\\\./~:]+$\"",
			}
		}

	default:
		return ResourceLocator_DirectiveValidationError{
			field:  "Directive",
			reason: "value is required",
		}

	}

	return nil
}

// ResourceLocator_DirectiveValidationError is the validation error returned by
// ResourceLocator_Directive.Validate if the designated constraints aren't met.
type ResourceLocator_DirectiveValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceLocator_DirectiveValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceLocator_DirectiveValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceLocator_DirectiveValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceLocator_DirectiveValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceLocator_DirectiveValidationError) ErrorName() string {
	return "ResourceLocator_DirectiveValidationError"
}

// Error satisfies the builtin error interface
func (e ResourceLocator_DirectiveValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceLocator_Directive.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceLocator_DirectiveValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceLocator_DirectiveValidationError{}

var _ResourceLocator_Directive_Entry_Pattern = regexp.MustCompile("^[0-9a-zA-Z_\\-\\./~:]+$")

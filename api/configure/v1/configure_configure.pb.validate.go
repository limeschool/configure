// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: configure_configure.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetConfigureRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetConfigureRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConfigureRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConfigureRequestMultiError, or nil if none found.
func (m *GetConfigureRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConfigureRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetServerId() <= 0 {
		err := GetConfigureRequestValidationError{
			field:  "ServerId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEnvId() <= 0 {
		err := GetConfigureRequestValidationError{
			field:  "EnvId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetConfigureRequestMultiError(errors)
	}

	return nil
}

// GetConfigureRequestMultiError is an error wrapping multiple validation
// errors returned by GetConfigureRequest.ValidateAll() if the designated
// constraints aren't met.
type GetConfigureRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConfigureRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConfigureRequestMultiError) AllErrors() []error { return m }

// GetConfigureRequestValidationError is the validation error returned by
// GetConfigureRequest.Validate if the designated constraints aren't met.
type GetConfigureRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConfigureRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConfigureRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConfigureRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConfigureRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConfigureRequestValidationError) ErrorName() string {
	return "GetConfigureRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetConfigureRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConfigureRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConfigureRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConfigureRequestValidationError{}

// Validate checks the field values on GetConfigureReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetConfigureReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConfigureReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConfigureReplyMultiError, or nil if none found.
func (m *GetConfigureReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConfigureReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ServerId

	// no validation rules for EnvId

	// no validation rules for Content

	// no validation rules for Description

	// no validation rules for Version

	// no validation rules for Format

	// no validation rules for CreatedAt

	if len(errors) > 0 {
		return GetConfigureReplyMultiError(errors)
	}

	return nil
}

// GetConfigureReplyMultiError is an error wrapping multiple validation errors
// returned by GetConfigureReply.ValidateAll() if the designated constraints
// aren't met.
type GetConfigureReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConfigureReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConfigureReplyMultiError) AllErrors() []error { return m }

// GetConfigureReplyValidationError is the validation error returned by
// GetConfigureReply.Validate if the designated constraints aren't met.
type GetConfigureReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConfigureReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConfigureReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConfigureReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConfigureReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConfigureReplyValidationError) ErrorName() string {
	return "GetConfigureReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetConfigureReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConfigureReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConfigureReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConfigureReplyValidationError{}

// Validate checks the field values on CompareConfigureRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CompareConfigureRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompareConfigureRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CompareConfigureRequestMultiError, or nil if none found.
func (m *CompareConfigureRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CompareConfigureRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetServerId() <= 0 {
		err := CompareConfigureRequestValidationError{
			field:  "ServerId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEnvId() <= 0 {
		err := CompareConfigureRequestValidationError{
			field:  "EnvId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CompareConfigureRequestMultiError(errors)
	}

	return nil
}

// CompareConfigureRequestMultiError is an error wrapping multiple validation
// errors returned by CompareConfigureRequest.ValidateAll() if the designated
// constraints aren't met.
type CompareConfigureRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompareConfigureRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompareConfigureRequestMultiError) AllErrors() []error { return m }

// CompareConfigureRequestValidationError is the validation error returned by
// CompareConfigureRequest.Validate if the designated constraints aren't met.
type CompareConfigureRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompareConfigureRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompareConfigureRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompareConfigureRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompareConfigureRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompareConfigureRequestValidationError) ErrorName() string {
	return "CompareConfigureRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CompareConfigureRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompareConfigureRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompareConfigureRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompareConfigureRequestValidationError{}

// Validate checks the field values on CompareConfigureInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CompareConfigureInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompareConfigureInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CompareConfigureInfoMultiError, or nil if none found.
func (m *CompareConfigureInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *CompareConfigureInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Key

	// no validation rules for Old

	// no validation rules for Cur

	if len(errors) > 0 {
		return CompareConfigureInfoMultiError(errors)
	}

	return nil
}

// CompareConfigureInfoMultiError is an error wrapping multiple validation
// errors returned by CompareConfigureInfo.ValidateAll() if the designated
// constraints aren't met.
type CompareConfigureInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompareConfigureInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompareConfigureInfoMultiError) AllErrors() []error { return m }

// CompareConfigureInfoValidationError is the validation error returned by
// CompareConfigureInfo.Validate if the designated constraints aren't met.
type CompareConfigureInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompareConfigureInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompareConfigureInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompareConfigureInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompareConfigureInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompareConfigureInfoValidationError) ErrorName() string {
	return "CompareConfigureInfoValidationError"
}

// Error satisfies the builtin error interface
func (e CompareConfigureInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompareConfigureInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompareConfigureInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompareConfigureInfoValidationError{}

// Validate checks the field values on CompareConfigureReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CompareConfigureReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompareConfigureReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CompareConfigureReplyMultiError, or nil if none found.
func (m *CompareConfigureReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CompareConfigureReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CompareConfigureReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CompareConfigureReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CompareConfigureReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CompareConfigureReplyMultiError(errors)
	}

	return nil
}

// CompareConfigureReplyMultiError is an error wrapping multiple validation
// errors returned by CompareConfigureReply.ValidateAll() if the designated
// constraints aren't met.
type CompareConfigureReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompareConfigureReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompareConfigureReplyMultiError) AllErrors() []error { return m }

// CompareConfigureReplyValidationError is the validation error returned by
// CompareConfigureReply.Validate if the designated constraints aren't met.
type CompareConfigureReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompareConfigureReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompareConfigureReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompareConfigureReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompareConfigureReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompareConfigureReplyValidationError) ErrorName() string {
	return "CompareConfigureReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CompareConfigureReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompareConfigureReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompareConfigureReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompareConfigureReplyValidationError{}

// Validate checks the field values on UpdateConfigureRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateConfigureRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateConfigureRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateConfigureRequestMultiError, or nil if none found.
func (m *UpdateConfigureRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateConfigureRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetServerId() <= 0 {
		err := UpdateConfigureRequestValidationError{
			field:  "ServerId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEnvId() <= 0 {
		err := UpdateConfigureRequestValidationError{
			field:  "EnvId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDescription()); l < 1 || l > 64 {
		err := UpdateConfigureRequestValidationError{
			field:  "Description",
			reason: "value length must be between 1 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateConfigureRequestMultiError(errors)
	}

	return nil
}

// UpdateConfigureRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateConfigureRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateConfigureRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateConfigureRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateConfigureRequestMultiError) AllErrors() []error { return m }

// UpdateConfigureRequestValidationError is the validation error returned by
// UpdateConfigureRequest.Validate if the designated constraints aren't met.
type UpdateConfigureRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateConfigureRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateConfigureRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateConfigureRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateConfigureRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateConfigureRequestValidationError) ErrorName() string {
	return "UpdateConfigureRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateConfigureRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateConfigureRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateConfigureRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateConfigureRequestValidationError{}

// Validate checks the field values on WatchConfigureRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WatchConfigureRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WatchConfigureRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WatchConfigureRequestMultiError, or nil if none found.
func (m *WatchConfigureRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WatchConfigureRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetServer()) < 1 {
		err := WatchConfigureRequestValidationError{
			field:  "Server",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetToken()) < 1 {
		err := WatchConfigureRequestValidationError{
			field:  "Token",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return WatchConfigureRequestMultiError(errors)
	}

	return nil
}

// WatchConfigureRequestMultiError is an error wrapping multiple validation
// errors returned by WatchConfigureRequest.ValidateAll() if the designated
// constraints aren't met.
type WatchConfigureRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WatchConfigureRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WatchConfigureRequestMultiError) AllErrors() []error { return m }

// WatchConfigureRequestValidationError is the validation error returned by
// WatchConfigureRequest.Validate if the designated constraints aren't met.
type WatchConfigureRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WatchConfigureRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WatchConfigureRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WatchConfigureRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WatchConfigureRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WatchConfigureRequestValidationError) ErrorName() string {
	return "WatchConfigureRequestValidationError"
}

// Error satisfies the builtin error interface
func (e WatchConfigureRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchConfigureRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WatchConfigureRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WatchConfigureRequestValidationError{}

// Validate checks the field values on WatchConfigureReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WatchConfigureReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WatchConfigureReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WatchConfigureReplyMultiError, or nil if none found.
func (m *WatchConfigureReply) ValidateAll() error {
	return m.validate(true)
}

func (m *WatchConfigureReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Content

	// no validation rules for Format

	if len(errors) > 0 {
		return WatchConfigureReplyMultiError(errors)
	}

	return nil
}

// WatchConfigureReplyMultiError is an error wrapping multiple validation
// errors returned by WatchConfigureReply.ValidateAll() if the designated
// constraints aren't met.
type WatchConfigureReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WatchConfigureReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WatchConfigureReplyMultiError) AllErrors() []error { return m }

// WatchConfigureReplyValidationError is the validation error returned by
// WatchConfigureReply.Validate if the designated constraints aren't met.
type WatchConfigureReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WatchConfigureReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WatchConfigureReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WatchConfigureReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WatchConfigureReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WatchConfigureReplyValidationError) ErrorName() string {
	return "WatchConfigureReplyValidationError"
}

// Error satisfies the builtin error interface
func (e WatchConfigureReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchConfigureReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WatchConfigureReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WatchConfigureReplyValidationError{}
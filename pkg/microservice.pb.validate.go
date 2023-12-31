// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: microservice.proto

package pkg

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

// Validate checks the field values on UploadFileRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UploadFileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadFileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadFileRequestMultiError, or nil if none found.
func (m *UploadFileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadFileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UploadFileRequestMultiError(errors)
	}

	return nil
}

// UploadFileRequestMultiError is an error wrapping multiple validation errors
// returned by UploadFileRequest.ValidateAll() if the designated constraints
// aren't met.
type UploadFileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadFileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadFileRequestMultiError) AllErrors() []error { return m }

// UploadFileRequestValidationError is the validation error returned by
// UploadFileRequest.Validate if the designated constraints aren't met.
type UploadFileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadFileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadFileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadFileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadFileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadFileRequestValidationError) ErrorName() string {
	return "UploadFileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadFileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadFileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadFileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadFileRequestValidationError{}

// Validate checks the field values on UploadFileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UploadFileResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadFileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadFileResponseMultiError, or nil if none found.
func (m *UploadFileResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadFileResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UploadFileResponseMultiError(errors)
	}

	return nil
}

// UploadFileResponseMultiError is an error wrapping multiple validation errors
// returned by UploadFileResponse.ValidateAll() if the designated constraints
// aren't met.
type UploadFileResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadFileResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadFileResponseMultiError) AllErrors() []error { return m }

// UploadFileResponseValidationError is the validation error returned by
// UploadFileResponse.Validate if the designated constraints aren't met.
type UploadFileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadFileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadFileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadFileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadFileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadFileResponseValidationError) ErrorName() string {
	return "UploadFileResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UploadFileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadFileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadFileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadFileResponseValidationError{}

// Validate checks the field values on DownloadFileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadFileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadFileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadFileRequestMultiError, or nil if none found.
func (m *DownloadFileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadFileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FileId

	if len(errors) > 0 {
		return DownloadFileRequestMultiError(errors)
	}

	return nil
}

// DownloadFileRequestMultiError is an error wrapping multiple validation
// errors returned by DownloadFileRequest.ValidateAll() if the designated
// constraints aren't met.
type DownloadFileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadFileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadFileRequestMultiError) AllErrors() []error { return m }

// DownloadFileRequestValidationError is the validation error returned by
// DownloadFileRequest.Validate if the designated constraints aren't met.
type DownloadFileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadFileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadFileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadFileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadFileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadFileRequestValidationError) ErrorName() string {
	return "DownloadFileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadFileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadFileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadFileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadFileRequestValidationError{}

// Validate checks the field values on MerkleProof with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MerkleProof) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MerkleProof with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MerkleProofMultiError, or
// nil if none found.
func (m *MerkleProof) ValidateAll() error {
	return m.validate(true)
}

func (m *MerkleProof) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Target

	if len(errors) > 0 {
		return MerkleProofMultiError(errors)
	}

	return nil
}

// MerkleProofMultiError is an error wrapping multiple validation errors
// returned by MerkleProof.ValidateAll() if the designated constraints aren't met.
type MerkleProofMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MerkleProofMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MerkleProofMultiError) AllErrors() []error { return m }

// MerkleProofValidationError is the validation error returned by
// MerkleProof.Validate if the designated constraints aren't met.
type MerkleProofValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MerkleProofValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MerkleProofValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MerkleProofValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MerkleProofValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MerkleProofValidationError) ErrorName() string { return "MerkleProofValidationError" }

// Error satisfies the builtin error interface
func (e MerkleProofValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMerkleProof.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MerkleProofValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MerkleProofValidationError{}

// Validate checks the field values on DownloadFileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadFileResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadFileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadFileResponseMultiError, or nil if none found.
func (m *DownloadFileResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadFileResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Content

	if all {
		switch v := interface{}(m.GetMerkleProof()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownloadFileResponseValidationError{
					field:  "MerkleProof",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownloadFileResponseValidationError{
					field:  "MerkleProof",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMerkleProof()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownloadFileResponseValidationError{
				field:  "MerkleProof",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DownloadFileResponseMultiError(errors)
	}

	return nil
}

// DownloadFileResponseMultiError is an error wrapping multiple validation
// errors returned by DownloadFileResponse.ValidateAll() if the designated
// constraints aren't met.
type DownloadFileResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadFileResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadFileResponseMultiError) AllErrors() []error { return m }

// DownloadFileResponseValidationError is the validation error returned by
// DownloadFileResponse.Validate if the designated constraints aren't met.
type DownloadFileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadFileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadFileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadFileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadFileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadFileResponseValidationError) ErrorName() string {
	return "DownloadFileResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadFileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadFileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadFileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadFileResponseValidationError{}

// Validate checks the field values on VerifyMerkleProofRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyMerkleProofRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyMerkleProofRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyMerkleProofRequestMultiError, or nil if none found.
func (m *VerifyMerkleProofRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyMerkleProofRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FileId

	// no validation rules for ClientMerkleRoot

	if len(errors) > 0 {
		return VerifyMerkleProofRequestMultiError(errors)
	}

	return nil
}

// VerifyMerkleProofRequestMultiError is an error wrapping multiple validation
// errors returned by VerifyMerkleProofRequest.ValidateAll() if the designated
// constraints aren't met.
type VerifyMerkleProofRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyMerkleProofRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyMerkleProofRequestMultiError) AllErrors() []error { return m }

// VerifyMerkleProofRequestValidationError is the validation error returned by
// VerifyMerkleProofRequest.Validate if the designated constraints aren't met.
type VerifyMerkleProofRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyMerkleProofRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyMerkleProofRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyMerkleProofRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyMerkleProofRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyMerkleProofRequestValidationError) ErrorName() string {
	return "VerifyMerkleProofRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyMerkleProofRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyMerkleProofRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyMerkleProofRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyMerkleProofRequestValidationError{}

// Validate checks the field values on VerifyMerkleProofResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyMerkleProofResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyMerkleProofResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyMerkleProofResponseMultiError, or nil if none found.
func (m *VerifyMerkleProofResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyMerkleProofResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsValid

	if len(errors) > 0 {
		return VerifyMerkleProofResponseMultiError(errors)
	}

	return nil
}

// VerifyMerkleProofResponseMultiError is an error wrapping multiple validation
// errors returned by VerifyMerkleProofResponse.ValidateAll() if the
// designated constraints aren't met.
type VerifyMerkleProofResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyMerkleProofResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyMerkleProofResponseMultiError) AllErrors() []error { return m }

// VerifyMerkleProofResponseValidationError is the validation error returned by
// VerifyMerkleProofResponse.Validate if the designated constraints aren't met.
type VerifyMerkleProofResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyMerkleProofResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyMerkleProofResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyMerkleProofResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyMerkleProofResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyMerkleProofResponseValidationError) ErrorName() string {
	return "VerifyMerkleProofResponseValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyMerkleProofResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyMerkleProofResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyMerkleProofResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyMerkleProofResponseValidationError{}

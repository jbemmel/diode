// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: diode/v1/site.proto

package diodepb

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

// Validate checks the field values on Site with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Site) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Site with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SiteMultiError, or nil if none found.
func (m *Site) ValidateAll() error {
	return m.validate(true)
}

func (m *Site) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Name != nil {

		if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 100 {
			err := SiteValidationError{
				field:  "Name",
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Slug != nil {

		if l := utf8.RuneCountInString(m.GetSlug()); l < 1 || l > 100 {
			err := SiteValidationError{
				field:  "Slug",
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if !_Site_Slug_Pattern.MatchString(m.GetSlug()) {
			err := SiteValidationError{
				field:  "Slug",
				reason: "value does not match regex pattern \"^[-a-zA-Z0-9_]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return SiteMultiError(errors)
	}

	return nil
}

// SiteMultiError is an error wrapping multiple validation errors returned by
// Site.ValidateAll() if the designated constraints aren't met.
type SiteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SiteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SiteMultiError) AllErrors() []error { return m }

// SiteValidationError is the validation error returned by Site.Validate if the
// designated constraints aren't met.
type SiteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SiteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SiteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SiteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SiteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SiteValidationError) ErrorName() string { return "SiteValidationError" }

// Error satisfies the builtin error interface
func (e SiteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSite.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SiteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SiteValidationError{}

var _Site_Slug_Pattern = regexp.MustCompile("^[-a-zA-Z0-9_]+$")

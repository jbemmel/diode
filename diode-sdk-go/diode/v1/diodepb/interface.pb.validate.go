// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: diode/v1/interface.proto

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

// Validate checks the field values on Interface with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Interface) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Interface with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InterfaceMultiError, or nil
// if none found.
func (m *Interface) ValidateAll() error {
	return m.validate(true)
}

func (m *Interface) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enabled

	if m.Device != nil {

		if m.GetDevice() == nil {
			err := InterfaceValidationError{
				field:  "Device",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if a := m.GetDevice(); a != nil {

		}

	}

	if m.Name != nil {

		if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 64 {
			err := InterfaceValidationError{
				field:  "Name",
				reason: "value length must be between 1 and 64 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Type != nil {

		if _, ok := _Interface_Type_InLookup[m.GetType()]; !ok {
			err := InterfaceValidationError{
				field:  "Type",
				reason: "value must be in list [virtual bridge lag 100base-fx 100base-lfx 100base-tx 100base-t1 1000base-t 1000base-x-gbic 1000base-x-sfp 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 50gbase-x-sfp56 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cxp 100gbase-x-cpak 100gbase-x-dsfp 100gbase-x-sfpdd 100gbase-x-qsfp28 100gbase-x-qsfpdd 200gbase-x-cfp2 200gbase-x-qsfp56 200gbase-x-qsfpdd 400gbase-x-cfp2 400gbase-x-qsfp112 400gbase-x-qsfpdd 400gbase-x-osfp 400gbase-x-osfp-rhs 400gbase-x-cdfp 400gbase-x-cfp8 800gbase-x-qsfpdd 800gbase-x-osfp 1000base-kx 10gbase-kr 10gbase-kx4 25gbase-kr 40gbase-kr4 50gbase-kr 100gbase-kp4 100gbase-kr2 100gbase-kr4 ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax ieee802.11ay ieee802.15.1 other-wireless gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 64gfc-qsfpp 128gfc-qsfp28 infiniband-sdr infiniband-ddr infiniband-qdr infiniband-fdr10 infiniband-fdr infiniband-edr infiniband-hdr infiniband-ndr infiniband-xdr t1 e1 t3 e3 xdsl docsis gpon xg-pon xgs-pon ng-pon2 epon 10g-epon cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus cisco-stackwise-80 cisco-stackwise-160 cisco-stackwise-320 cisco-stackwise-480 cisco-stackwise-1t juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Mtu != nil {

		if val := m.GetMtu(); val < 1 || val > 65536 {
			err := InterfaceValidationError{
				field:  "Mtu",
				reason: "value must be inside range [1, 65536]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.MacAddress != nil {
		// no validation rules for MacAddress
	}

	if m.MgmtOnly != nil {
		// no validation rules for MgmtOnly
	}

	if len(errors) > 0 {
		return InterfaceMultiError(errors)
	}

	return nil
}

// InterfaceMultiError is an error wrapping multiple validation errors returned
// by Interface.ValidateAll() if the designated constraints aren't met.
type InterfaceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InterfaceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InterfaceMultiError) AllErrors() []error { return m }

// InterfaceValidationError is the validation error returned by
// Interface.Validate if the designated constraints aren't met.
type InterfaceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InterfaceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InterfaceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InterfaceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InterfaceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InterfaceValidationError) ErrorName() string { return "InterfaceValidationError" }

// Error satisfies the builtin error interface
func (e InterfaceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInterface.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InterfaceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InterfaceValidationError{}

var _Interface_Type_InLookup = map[string]struct{}{
	"virtual":                 {},
	"bridge":                  {},
	"lag":                     {},
	"100base-fx":              {},
	"100base-lfx":             {},
	"100base-tx":              {},
	"100base-t1":              {},
	"1000base-t":              {},
	"1000base-x-gbic":         {},
	"1000base-x-sfp":          {},
	"2.5gbase-t":              {},
	"5gbase-t":                {},
	"10gbase-t":               {},
	"10gbase-cx4":             {},
	"10gbase-x-sfpp":          {},
	"10gbase-x-xfp":           {},
	"10gbase-x-xenpak":        {},
	"10gbase-x-x2":            {},
	"25gbase-x-sfp28":         {},
	"50gbase-x-sfp56":         {},
	"40gbase-x-qsfpp":         {},
	"50gbase-x-sfp28":         {},
	"100gbase-x-cfp":          {},
	"100gbase-x-cfp2":         {},
	"100gbase-x-cfp4":         {},
	"100gbase-x-cxp":          {},
	"100gbase-x-cpak":         {},
	"100gbase-x-dsfp":         {},
	"100gbase-x-sfpdd":        {},
	"100gbase-x-qsfp28":       {},
	"100gbase-x-qsfpdd":       {},
	"200gbase-x-cfp2":         {},
	"200gbase-x-qsfp56":       {},
	"200gbase-x-qsfpdd":       {},
	"400gbase-x-cfp2":         {},
	"400gbase-x-qsfp112":      {},
	"400gbase-x-qsfpdd":       {},
	"400gbase-x-osfp":         {},
	"400gbase-x-osfp-rhs":     {},
	"400gbase-x-cdfp":         {},
	"400gbase-x-cfp8":         {},
	"800gbase-x-qsfpdd":       {},
	"800gbase-x-osfp":         {},
	"1000base-kx":             {},
	"10gbase-kr":              {},
	"10gbase-kx4":             {},
	"25gbase-kr":              {},
	"40gbase-kr4":             {},
	"50gbase-kr":              {},
	"100gbase-kp4":            {},
	"100gbase-kr2":            {},
	"100gbase-kr4":            {},
	"ieee802.11a":             {},
	"ieee802.11g":             {},
	"ieee802.11n":             {},
	"ieee802.11ac":            {},
	"ieee802.11ad":            {},
	"ieee802.11ax":            {},
	"ieee802.11ay":            {},
	"ieee802.15.1":            {},
	"other-wireless":          {},
	"gsm":                     {},
	"cdma":                    {},
	"lte":                     {},
	"sonet-oc3":               {},
	"sonet-oc12":              {},
	"sonet-oc48":              {},
	"sonet-oc192":             {},
	"sonet-oc768":             {},
	"sonet-oc1920":            {},
	"sonet-oc3840":            {},
	"1gfc-sfp":                {},
	"2gfc-sfp":                {},
	"4gfc-sfp":                {},
	"8gfc-sfpp":               {},
	"16gfc-sfpp":              {},
	"32gfc-sfp28":             {},
	"64gfc-qsfpp":             {},
	"128gfc-qsfp28":           {},
	"infiniband-sdr":          {},
	"infiniband-ddr":          {},
	"infiniband-qdr":          {},
	"infiniband-fdr10":        {},
	"infiniband-fdr":          {},
	"infiniband-edr":          {},
	"infiniband-hdr":          {},
	"infiniband-ndr":          {},
	"infiniband-xdr":          {},
	"t1":                      {},
	"e1":                      {},
	"t3":                      {},
	"e3":                      {},
	"xdsl":                    {},
	"docsis":                  {},
	"gpon":                    {},
	"xg-pon":                  {},
	"xgs-pon":                 {},
	"ng-pon2":                 {},
	"epon":                    {},
	"10g-epon":                {},
	"cisco-stackwise":         {},
	"cisco-stackwise-plus":    {},
	"cisco-flexstack":         {},
	"cisco-flexstack-plus":    {},
	"cisco-stackwise-80":      {},
	"cisco-stackwise-160":     {},
	"cisco-stackwise-320":     {},
	"cisco-stackwise-480":     {},
	"cisco-stackwise-1t":      {},
	"juniper-vcp":             {},
	"extreme-summitstack":     {},
	"extreme-summitstack-128": {},
	"extreme-summitstack-256": {},
	"extreme-summitstack-512": {},
	"other":                   {},
}

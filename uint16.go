package vl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/speedyhoon/frm"
	"github.com/speedyhoon/utl"
)

// Uint16 validates inp as an unsigned 16-bit integer.
func Uint16(f *frm.Field, inp ...string) {
	if !parseUint16(f, inp...) {
		return
	}
	value := f.Uint16()
	if f.Min != nil && value < f.Min.(uint16) || f.Max != nil && value > f.Max.(uint16) {
		f.Err = fmt.Sprintf("Must be between %d and %d.", f.Min, f.Max)
		return
	}

	var step uint16
	if f.Step == 0 {
		step = 1
	} else {
		step = uint16(f.Step)
	}

	if value%step != 0 {
		below := value - value%step
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %d and %d.", below, below+step)
		return
	}
}

// Uint16List validates inp as a slice of unsigned 16-bit integers.
func Uint16List(f *frm.Field, inp ...string) {
	if len(inp) < f.MinLen {
		f.Err = fmt.Sprintf("Not enough items selected. At least %d item%s required.", f.MinLen, utl.Plural(len(inp), " is", "s are"))
		return
	}

	var list []uint16

	for _, str := range inp {
		Uint(f, str)
		if f.Err != "" {
			return
		}

		value := f.Uint16()

		// Check if this value isn't already in the list.
		for _, num := range list {
			if value == num {
				f.Err = "Duplicate values found in the list."
				return
			}
		}

		list = append(list, value)
	}

	f.Value = list
}

// Uint16Req enforces an unsigned 16-bit integer to be required.
func Uint16Req(f *frm.Field, inp ...string) {
	f.Required = true
	Uint16(f, inp...)
}

// Uint16Opt unsigned 16-bit integer option slice.
func Uint16Opt(f *frm.Field, inp ...string) {
	if !parseUint16(f, inp...) || len(f.Options) < 1 {
		return
	}

	var ok bool
	for _, option := range f.Options {
		ok = f.Value == option.Value
		if ok {
			break
		}
	}
	if !ok {
		f.Err = "Please select from one of the options."
	}
}

// parseUint16 returns false upon validation failure.
func parseUint16(f *frm.Field, inp ...string) bool {
	u, err := strconv.ParseUint(strings.TrimSpace(inp[0]), 10, 16)
	if err != nil {
		// Return error if input string failed to convert.
		f.Err = err.Error()
		return false
	}

	f.Value = u
	return f.Required && f.Uint16() != 0
}

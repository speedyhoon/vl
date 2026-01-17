package vl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/speedyhoon/frm"
	"github.com/speedyhoon/utl"
)

// Uint8 validates inp as an unsigned 8-bit integer.
func Uint8(f *frm.Field, inp ...string) {
	if !parseUint8(f, inp...) {
		return
	}
	value := f.Uint8()
	if f.Min != nil && value < f.Min.(uint8) || f.Max != nil && value > f.Max.(uint8) {
		f.Err = fmt.Sprintf("Must be between %d and %d.", f.Min, f.Max)
		return
	}

	var step uint8
	if f.Step == 0 {
		step = 1
	} else {
		step = uint8(f.Step)
	}

	if value%step != 0 {
		below := value - value%step
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %d and %d.", below, below+step)
		return
	}
}

// Uint8List validates inp as a slice of unsigned 8-bit integers.
func Uint8List(f *frm.Field, inp ...string) {
	if len(inp) < f.MinLen {
		f.Err = fmt.Sprintf("Not enough items selected. At least %d item%s required.", f.MinLen, utl.Plural(len(inp), " is", "s are"))
		return
	}

	var list []uint8

	for _, str := range inp {
		Uint(f, str)
		if f.Err != "" {
			return
		}

		value := f.Uint8()

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

// Uint8Req enforces an unsigned 8-bit integer to be required.
func Uint8Req(f *frm.Field, inp ...string) {
	f.Required = true
	Uint8(f, inp...)
}

// Uint8Opt unsigned 8-bit integer option slice.
func Uint8Opt(f *frm.Field, inp ...string) {
	if !parseUint8(f, inp...) || len(f.Options) < 1 {
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

// parseUint8 returns false upon validation failure.
func parseUint8(f *frm.Field, inp ...string) bool {
	u, err := strconv.ParseUint(strings.TrimSpace(inp[0]), 10, 8)
	if err != nil {
		// Return error if input string failed to convert.
		f.Err = err.Error()
		return false
	}

	f.Value = u
	return f.Required && f.Uint8() != 0
}

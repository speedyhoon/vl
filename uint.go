package vl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/speedyhoon/frm"
	"github.com/speedyhoon/utl"
)

// Uint validates inp as an unsigned integer.
func Uint(f *frm.Field, inp ...string) {
	if !parseUint(f, inp...) {
		return
	}
	value := f.Uint()
	if f.Min != nil && value < f.Min.(uint) || f.Max != nil && value > f.Max.(uint) {
		f.Err = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	var step uint
	if f.Step == 0 {
		step = 1
	} else {
		step = uint(f.Step)
	}

	if value%step != 0 {
		below := value - value%step
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %d and %d.", below, below+step)
		return
	}
}

// UintList validates inp as a slice of unsigned integers.
func UintList(f *frm.Field, inp ...string) {
	if len(inp) < f.MinLen {
		f.Err = fmt.Sprintf("Not enough items selected. At least %v item%s required.", f.MinLen, utl.Plural(len(inp), " is", "s are"))
		return
	}

	var list []uint

	for _, str := range inp {
		Uint(f, str)
		if f.Err != "" {
			return
		}

		value := f.Uint()

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

// UintReq enforces an unsigned integer to be required.
func UintReq(f *frm.Field, inp ...string) {
	f.Required = true
	Uint(f, inp...)
}

// UintOpt unsigned integer option slice.
func UintOpt(f *frm.Field, inp ...string) {
	if !parseUint(f, inp...) || len(f.Options) < 1 {
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

// parseUint returns false upon validation failure.
func parseUint(f *frm.Field, inp ...string) bool {
	u, err := strconv.ParseUint(strings.TrimSpace(inp[0]), 10, sysArch)
	if err != nil {
		//Return error if input string failed to convert.
		f.Err = err.Error()
		return false
	}

	f.Value = uint(u)
	return f.Required && f.Uint() != 0
}

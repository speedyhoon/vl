package vl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/speedyhoon/frm"
	"github.com/speedyhoon/utl"
)

// Uint32 validates inp as an unsigned 32-bit integer
func Uint32(f *frm.Field, inp ...string) {
	if !parseUint32(f, inp...) {
		return
	}
	value := f.Uint32()
	if f.Min != nil && value < f.Min.(uint32) || f.Max != nil && value > f.Max.(uint32) {
		f.Err = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	var step uint32
	if f.Step == 0 {
		step = 1
	} else {
		step = uint32(f.Step)
	}

	if value%step != 0 {
		below := value - value%step
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %d and %d.", below, below+step)
		return
	}
}

// Uint32List validates inp as a slice of unsigned 32-bit integers
func Uint32List(f *frm.Field, inp ...string) {
	if len(inp) < f.MinLen {
		f.Err = fmt.Sprintf("Not enough items selected. At least %v item%s required.", f.MinLen, utl.Plural(len(inp), " is", "s are"))
		return
	}

	var list []uint32

	for _, str := range inp {
		Uint(f, str)
		if f.Err != "" {
			return
		}

		value := f.Uint32()

		//check if this value isn't already in the list
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

// Uint32Req enforces an unsigned 32-bit integer to be required
func Uint32Req(f *frm.Field, inp ...string) {
	f.Required = true
	Uint32(f, inp...)
}

// Uint32Opt unsigned 32-bit integer option slice
func Uint32Opt(f *frm.Field, inp ...string) {
	if !parseUint32(f, inp...) || len(f.Options) < 1 {
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

// parseUint32 returns false upon validation failure
func parseUint32(f *frm.Field, inp ...string) bool {
	u, err := strconv.ParseUint(strings.TrimSpace(inp[0]), 10, 32)
	if err != nil {
		//Return error if input string failed to convert.
		f.Err = err.Error()
		return false
	}

	f.Value = u
	return f.Required && f.Uint32() != 0
}

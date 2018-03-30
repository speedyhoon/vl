package v8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/speedyhoon/forms"
	"github.com/speedyhoon/utl"
)

func Uint(f *forms.Field, inp ...string) {
	if !parseUint(f, inp...) {
		return
	}
	if f.ValueUint < uint(f.Min) || f.ValueUint > uint(f.Max) {
		f.Error = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	var step uint
	if f.Step == 0 {
		step = 1
	}else{
		step = uint(f.Step)
	}

	if f.ValueUint%uint(f.Step) != 0 {
		below := f.ValueUint - f.ValueUint % step
		f.Error = fmt.Sprintf("Please enter a valid value. The two nearest values are %d and %d.", below, below + step)
		return
	}
}

func UintList(f *forms.Field, inp ...string) {
	if len(inp) < f.MinLen {
		f.Error = fmt.Sprintf("Not enough items selected. At least %v item%s required.", f.MinLen, util.Plural(len(inp), " is", "s are"))
		return
	}

	check := make(map[uint]bool, len(inp))
	var list []uint

	for _, in := range inp {
		Uint(f, in)
		if f.Error != "" {
			return
		}

		_, ok := check[f.ValueUint]
		if ok {
			f.Error = "Duplicate values found in list."
			return
		}
		check[f.ValueUint] = true
		list = append(list, f.ValueUint)
	}

	f.ValueUintSlice = list
}

//Required unsigned integer
func UintReq(f *forms.Field, inp ...string) {
	f.Required = true
	Uint(f, inp...)
}

//unsigned integer option
func UintOpt(f *forms.Field, inp ...string) {
	if !parseUint(f, inp...) {
		return
	}

	var found bool
	for _, option := range f.Options {
		if f.Value == option.Value {
			found = true
			break
		}
	}
	if !found {
		f.Error = "Please select from one of the options."
	}
}

//parseUint returns false upon validation failure
func parseUint(f *forms.Field, inp ...string) bool {
	f.Value = strings.TrimSpace(inp[0])
	u, err := strconv.ParseUint(f.Value, 10, sysArch)
	if err != nil {
		//Return error if input string failed to convert.
		f.Error = err.Error()
		return false
	}

	return f.Required && uint(u) != 0
}
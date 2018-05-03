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
	value := f.Uint()
	if value < uint(f.Min) || value > uint(f.Max) {
		f.Err = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	var step uint
	if f.Step == 0 {
		step = 1
	}else{
		step = uint(f.Step)
	}

	if value % step != 0 {
		below := value - value % step
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %d and %d.", below, below + step)
		return
	}
}

func UintList(f *forms.Field, inp ...string) {
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

//Required unsigned integer
func UintReq(f *forms.Field, inp ...string) {
	f.Required = true
	Uint(f, inp...)
}

//unsigned integer option
func UintOpt(f *forms.Field, inp ...string) {
	if !parseUint(f, inp...) || len(f.Options) < 1 {
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
		f.Err = "Please select from one of the options."
	}
}

//parseUint returns false upon validation failure
func parseUint(f *forms.Field, inp ...string) bool {
	f.Value = strings.TrimSpace(inp[0])
	u, err := strconv.ParseUint(f.Str(), 10, sysArch)
	if err != nil {
		//Return error if input string failed to convert.
		f.Err = err.Error()
		return false
	}

	f.Value = uint(u)
	return f.Required && f.Uint() != 0
}
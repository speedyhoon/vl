package v8

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/speedyhoon/forms"
	"github.com/speedyhoon/util"
)

const maxLen int = 64

func UintList(f *forms.Field, inp ...string) {
	if len(inp) < f.MinLen {
		f.Error = fmt.Sprintf("Not enough items selected. At least %v item%s needed.", f.MinLen, util.Plural(len(inp), " is", "s are"))
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

//UintBasic returns false upon validation failure
func UintBasic(f *forms.Field, inp ...string) bool {
	f.Value = inp[0]
	u, err := strconv.ParseUint(strings.TrimSpace(f.Value), 10, sysArch)
	if err != nil {
		//Return error if input string failed to convert.
		f.Error = err.Error()
		return false
	}

	return f.Required && uint(u) != 0
}

func Uint(f *forms.Field, inp ...string) {
	if !UintBasic(f, inp...) {
		return
	}
	if f.ValueUint < uint(f.Min) || f.ValueUint > uint(f.Max) {
		f.Error = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	if f.Step == 0 {
		f.Step = 1
	}
	if f.ValueUint%uint(f.Step) != 0 {
		below := f.ValueUint - f.ValueUint%uint(f.Step)
		f.Error = fmt.Sprintf("Please enter a valid value. The two nearest values are %d and %d.", below, below+uint(f.Step))
		return
	}
}

func UintReq(f *forms.Field, inp ...string) {
	f.Required = true
	Uint(f, inp...)
}

func UintOpt(f *forms.Field, inp ...string) {
	if !UintBasic(f, inp...) {
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
		f.Error = "Please select from one of the options"
	}
}

func Float32(f *forms.Field, inp ...string) {
	f64, err := strconv.ParseFloat(strings.TrimSpace(inp[0]), 32)
	if err != nil {
		//Return error if input string failed to convert.
		f.Error = err.Error()
		return
	}
	num := float32(f64)

	if !f.Required && num == 0 {
		//f.ValueFloat32 is zero by default so assigning zero isn't required
		return
	}
	if num < float32(f.Min) || num > float32(f.Max) {
		f.Error = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	if rem := toFixed(math.Mod(f64, float64(f.Step)), 6); rem != 0 {
		f.Error = fmt.Sprintf("Please enter a valid value. The two nearest values are %v and %v.", num-rem, num-rem+f.Step)
		return
	}
	f.Value = fmt.Sprintf("%v", num)
	f.ValueFloat32 = num
}

func toFixed(num, precision float64) float32 {
	output := math.Pow(10, precision)
	return float32(int(num * output)) / float32(output)
}

func Str(f *forms.Field, inp ...string) {
	f.Value = strings.TrimSpace(inp[0])

	//Check value matches regex
	if f.Regex != nil && !f.Regex.MatchString(f.Value) {
		f.Error = "Failed pattern"
		return
	}

	if f.MinLen == 0 && f.Required {
		f.MinLen = 1
	}
	length := len(f.Value)
	if length < f.MinLen {
		f.Error = fmt.Sprintf("Please lengthen this text to %d characters or more (you are currently using %d character%v).", f.MinLen, length, util.Plural(length, "", ""))
		return
	}

	if f.MaxLen == 0 {
		f.MaxLen = maxLen
	}
	if length > f.MaxLen {
		//Truncate string instead of raising an error
		f.Value = f.Value[:f.MaxLen]
	}

	//Check value matches one of the options (optional).
	/*if len(f.Options) > 0 {
		matched := false
		for _, option := range f.Options {
			matched = option.Value == f.Value
			if matched {
				break
			}
		}
		if !matched {
			f.Error = "Value doesn't match any of the options"
			return
		}
	}*/
}

func StrReq(f *forms.Field, inp ...string) {
	f.Required = true
	Str(f, inp...)
}

func Regex(f *forms.Field, inp ...string) {
	f.Value = strings.TrimSpace(inp[0])
	if !f.Regex.MatchString(f.Value) {
		f.Error = "ID supplied is incorrect."
	}
}

func RegexReq(f *forms.Field, inp ...string) {
	f.Required = true
	if inp[0] != "" {
		Str(f, inp...)
	}
	f.Error = "Empty ID supplied."
}

func Bool(f *forms.Field, inp ...string) {
	f.Checked = len(strings.TrimSpace(inp[0])) >= 1
	if f.Required && !f.Checked {
		f.Error = "Please check this field."
	}
}

func FileReq(f *forms.Field, inp ...string) {
	//TODO add validation checks
	//maxlength < 2MB?
	//Unmarshal??
	//Return as interface{}??
}


package v8

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/speedyhoon/forms"
	"github.com/speedyhoon/utl"
)

const maxLen int = 64

func UintList(f *forms.Field, inp ...string) {
	if len(inp) < f.MinLen {
		f.Err = fmt.Sprintf("Not enough items selected. At least %v item%s required.", f.MinLen, util.Plural(len(inp), " is", "s are"))
		return
	}

	check := make(map[uint]bool, len(inp))
	var list []uint

	for _, in := range inp {
		Uint(f, in)
		if f.Err != "" {
			return
		}

		value := f.Uint()
		_, ok := check[value]
		if ok {
			f.Err = "Duplicate values found in list."
			return
		}
		check[value] = true
		list = append(list, value)
	}

	f.Value = list
}

//UintBasic returns false upon validation failure
func UintBasic(f *forms.Field, inp ...string) bool {
	value := inp[0]
	u64, err := strconv.ParseUint(strings.TrimSpace(value), 10, sysArch)
	if err != nil {
		//Return error if input string failed to convert.
		f.Err = err.Error()
		return false
	}

	f.Value = uint(u64)
	return f.Required && uint(u64) != 0
}

func Uint(f *forms.Field, inp ...string) {
	if !UintBasic(f, inp...) {
		return
	}

	value := f.Uint()
	if value < uint(f.Min) || value > uint(f.Max) {
		f.Err = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	if f.Step == 0 {
		f.Step = 1
	}
	if value%uint(f.Step) != 0 {
		below := value - value%uint(f.Step)
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %d and %d.", below, below+uint(f.Step))
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
		f.Err = "Please select from one of the options."
	}
}

func Float32(f *forms.Field, inp ...string) {
	f64, err := strconv.ParseFloat(strings.TrimSpace(inp[0]), 32)
	if err != nil {
		//Return error if input string failed to convert.
		f.Err = err.Error()
		return
	}
	num := float32(f64)

	if !f.Required && num == 0 {
		//f.ValueFloat32 is zero by default so assigning zero isn't required
		return
	}
	if num < float32(f.Min) || num > float32(f.Max) {
		f.Err = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	if rem := toFixed(math.Mod(f64, float64(f.Step)), 6); rem != 0 {
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %v and %v.", num-rem, num-rem+f.Step)
		return
	}
	f.Value = num
}

func toFixed(num, precision float64) float32 {
	output := math.Pow(10, precision)
	return float32(int(num * output)) / float32(output)
}

func Str(f *forms.Field, inp ...string) {
	value := strings.TrimSpace(inp[0])
	f.Value = value

	//Check value matches regex
	if f.Regex != nil && !f.Regex.MatchString(value) {
		f.Err = "Failed pattern."
		return
	}

	if f.MinLen == 0 && f.Required {
		f.MinLen = 1
	}
	length := len(value)
	if length < f.MinLen {
		f.Err = fmt.Sprintf("Please lengthen this text to %d characters or more (you are currently using %d character%v).", f.MinLen, length, util.Plural(length, "", ""))
		return
	}

	if f.MaxLen == 0 {
		f.MaxLen = maxLen
	}
	if length > f.MaxLen {
		//Truncate string instead of raising an error
		value = value[:f.MaxLen]
	}
	f.Value = value

	//Check value matches one of the options (optional).
	/*if len(f.Options) > 0 {
		matched := false
		for _, option := range f.Options {
			matched = option.Value == value
			if matched {
				break
			}
		}
		if !matched {
			f.Err = "Value doesn't match any of the options."
			return
		}
	}*/
}

func StrReq(f *forms.Field, inp ...string) {
	f.Required = true
	Str(f, inp...)
}

func Regex(f *forms.Field, inp ...string) {
	value := strings.TrimSpace(inp[0])
	f.Value = value
	if f.Required && value == "" {
		f.Err = "Empty ID supplied."
		return
	}
	if !f.Regex.MatchString(value) {
		f.Err = "ID supplied is incorrect."
	}
}

func RegexReq(f *forms.Field, inp ...string) {
	f.Required = true
	Str(f, inp...)
}

func Bool(f *forms.Field, inp ...string) {
	f.Value = len(strings.TrimSpace(inp[0])) >= 1
	if f.Required && !f.Checked() {
		f.Err = "Please check this field."
	}
}

func FileReq(f *forms.Field, inp ...string) {
	//TODO add validation checks
	//maxlength < 2MB?
	//Unmarshal??
	//Return as interface{}??
}

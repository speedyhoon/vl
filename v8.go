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

func Str(f *forms.Field, inp ...string) {
	f.Value = strings.TrimSpace(inp[0])

	//Check value matches regex
	if f.Regex != nil && !f.Regex.MatchString(f.Value) {
		f.Error = "Failed pattern."
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
			f.Error = "Value doesn't match any of the options."
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

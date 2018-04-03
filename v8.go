package v8

import (
	"fmt"
	"strings"

	"github.com/speedyhoon/forms"
	"github.com/speedyhoon/utl"
)

const maxLen int = 64

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
}

//Check value matches one of the options (optional).
func StrOpt(f *forms.Field, inp ...string) {
	Str(f, inp...)

	if f.Err != "" || len(f.Options) < 1 {
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
		f.Err = "Value doesn't match any of the options."
		return
	}
}

func StrReq(f *forms.Field, inp ...string) {
	f.Required = true
	Str(f, inp...)
}

func Regex(f *forms.Field, inp ...string) {
	f.Value = strings.TrimSpace(inp[0])
	if f.Required && f.Str() == "" {
		f.Err = "Empty ID supplied."
		return
	}
	if !f.Regex.MatchString(f.Str()) {
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

package vl

import (
	"fmt"
	"time"

	"github.com/speedyhoon/frm"
)

//DateTime validates inp as a time.Time input
func DateTime(f *frm.Field, inp ...string) {
	value, err := time.Parse(f.Placeholder, inp[0])
	f.Value = value
	if err != nil {
		f.Err = err.Error()
		return
	}

	if f.Min != nil && f.Min.(time.Time).After(value) || f.Max != nil && value.After(f.Max.(time.Time)) {
		f.Err = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}
}

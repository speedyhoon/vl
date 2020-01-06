package vl

import (
	"time"

	"github.com/speedyhoon/frm"
)

//DateTime validates inp as a time.Time input
func DateTime(f *frm.Field, inp ...string) {
	var err error
	f.Value, err = time.Parse(f.Placeholder, inp[0])
	if err != nil {
		f.Err = err.Error()
	}
}

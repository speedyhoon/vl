package v8

import (
	"github.com/speedyhoon/forms"
	"time"
)

func DateTime(f *forms.Field, inp ...string) {
	var err error
	f.Value, err = time.Parse(f.Placeholder, inp[0])
	if err != nil{
		f.Err = err.Error()
	}
}

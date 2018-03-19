package v8

import (
	"net/http"
	"net/url"
	"strings"
	"github.com/speedyhoon/forms"
)

func IsValidRequest(r *http.Request, f forms.Form) (forms.Form, bool) {
	var err error
	var u *url.URL

	isGet := r.Method == "GET"
	if isGet {
		u, err = url.Parse(r.RequestURI)
	} else {
		err = r.ParseForm()
	}

	if err != nil {
		return f, false
	}

	if isGet {
		return IsValid(u.Query(), f)
	}
	return IsValid(r.Form, f)
}

//Is it worth while to auto add failed forms to session so it doesn't have to be done in each http handler?
func IsValid(urlValues url.Values, f forms.Form) (forms.Form, bool) {
	if len(urlValues) == 0 {
		return f, false
	}
	//Process the post request as normal if len(urlValues) >= len(fields).
	var fieldValue []string
	var ok bool
	isValid := true
	for i, field := range f.Fields {
		/*// Output warning if validation function is not set for this field in the submitted form.
		if debug && field.v8 == nil {
			field.Error = "No v8 function setup!"
			warn.Println("No v8 function setup! for", field.name)
			continue
		}*/
		fieldValue, ok = urlValues[field.Name]

		//if fieldValue is empty and field is required
		if !ok || len(fieldValue) == 0 || len(fieldValue) == 1 && strings.TrimSpace(fieldValue[0]) == "" {
			if field.Required {
				f.Fields[i].Error = "Please fill in this field."
			}
			//else if field is not required & its contents is empty - don't validate
		} else {
			//Otherwise validate user input
			field.V8(&f.Fields[i], fieldValue...)
		}

		//Set the first field with failed validation to have focus onscreen
		if f.Fields[i].Error != "" && isValid {
			f.Fields[i].AutoFocus = true
			isValid = false
		}
	}
	return f, isValid
}

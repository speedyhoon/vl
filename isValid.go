package vl

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/speedyhoon/frm"
)

//IsValidRequest gathers the form submitted from GET and POST requests and then calls IsValid()
func IsValidRequest(r *http.Request, fields []frm.Field) ([]frm.Field, bool) {
	var err error
	var u *url.URL

	isGet := r.Method == "GET"
	if isGet {
		u, err = url.Parse(r.RequestURI)
	} else {
		err = r.ParseForm()
	}

	if err != nil {
		return fields, false
	}

	if isGet {
		return IsValid(u.Query(), fields)
	}
	return IsValid(r.Form, fields)
}

//IsValid loops through each form field and validates with a function from v8
func IsValid(urlValues url.Values, fields []frm.Field) ([]frm.Field, bool) {
	if len(urlValues) == 0 {
		//TODO Is it worth while to auto add failed forms to session so it doesn't have to be done in each http handler?
		return fields, false
	}
	//Process the post request as normal if len(urlValues) >= len(fields).
	var fieldValue []string
	var ok bool
	isValid := true
	for i := range fields {
		/*// Output warning if validation function is not set for this field in the submitted form.
		if debug && fields[i].V8 == nil {
			fields[i].Err = "No validation function setup for " + fields[i].Name
			continue
		}*/
		fieldValue, ok = urlValues[fields[i].Name]

		//if fieldValue is empty and field is required
		if !ok || len(fieldValue) == 0 || len(fieldValue) == 1 && strings.TrimSpace(fieldValue[0]) == "" {
			if fields[i].Required {
				fields[i].Err = "Please fill in this field."
			} else {
				//else if field isn't required & its contents is empty, then don't validate
				continue
			}
		} else {
			//Otherwise validate user input
			fields[i].V8(&fields[i], fieldValue...)
		}

		//Set the first field with failed validation to have focus onscreen
		if fields[i].Err != "" && isValid {
			fields[i].Focus = true
			isValid = false
		}
	}
	return fields, isValid
}

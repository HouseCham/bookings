package forms

import (
	"net/http"
	"net/url"
)

//? Creates a custom form struct and it embeds a url.values object
type Form struct {
	url.Values
	Errors errors
}

//? Initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//? Checks if  form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x != ""{
		return true
	}
	return false
}
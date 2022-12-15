package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// create form struct
type Form struct {
	url.Values
	Errors errors
}

// create new function to init form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// implement required method to check if field is not blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// implement MaxLength method
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

// implement permitted values method to check a specifi field in the form
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// implements a valid method to check if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

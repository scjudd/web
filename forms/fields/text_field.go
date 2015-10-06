package fields

import (
	"fmt"
)

type TextField struct {
	Name     string
	Value    string
	Required bool
	errors   []error
}

func (f *TextField) HTML() string {
	return fmt.Sprintf("<input type=\"text\" name=\"%s\" value=\"%s\">", f.Name, f.Value)
}

func (f *TextField) GetName() string {
	return f.Name
}

func (f *TextField) GetValue() string {
	return f.Value
}

func (f *TextField) SetValue(value string) {
	f.Value = value
}

func (f *TextField) AddError(errs ...error) {
	for _, err := range errs {
		f.errors = append(f.errors, err)
	}
}

func (f *TextField) Valid() (bool, string) {
	if f.Required && f.Value == "" {
		return false, "This field is required."
	}
	if len(f.errors) > 0 {
		return false, f.errors[0].Error()
	}
	return true, ""
}

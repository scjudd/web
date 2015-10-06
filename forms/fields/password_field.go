package fields

import (
	"fmt"
)

type PasswordField struct {
	Name     string
	Value    string
	Required bool
	errors   []error
}

func (f *PasswordField) HTML() string {
	return fmt.Sprintf("<input type=\"password\" name=\"%s\" value=\"%s\">", f.Name, f.Value)
}

func (f *PasswordField) GetName() string {
	return f.Name
}

func (f *PasswordField) GetValue() string {
	return f.Value
}

func (f *PasswordField) SetValue(value string) {
	f.Value = value
}

func (f *PasswordField) AddError(errs ...error) {
	for _, err := range errs {
		f.errors = append(f.errors, err)
	}
}

func (f *PasswordField) Valid() (bool, string) {
	if f.Required && f.Value == "" {
		return false, "This field is required."
	}
	if len(f.errors) > 0 {
		return false, f.errors[0].Error()
	}
	return true, ""
}

package forms

import (
	"fmt"
	"html/template"
	"strings"
)

type Field interface {
	Valid() (bool, string)
	Type() string
	GetValue() string
	SetValue(string)
	GetClasses() []string
	AddClass(string)
}

func FieldHTML(f Field, name string) template.HTML {
	return template.HTML(fmt.Sprintf(
		"<input type=\"%s\" name=\"%s\" class=\"%s\" value=\"%s\">",
		f.Type(), name, strings.Join(f.GetClasses(), " "), f.GetValue()))
}

type BaseField struct {
	Value   string
	Classes []string
}

func (f *BaseField) SetValue(value string) {
	f.Value = value
}

func (f *BaseField) GetValue() string {
	return f.Value
}

func (f *BaseField) GetClasses() []string {
	return f.Classes
}

func (f *BaseField) AddClass(class string) {
	f.Classes = append(f.Classes, class)
}

type TextField struct {
	BaseField
	Required bool
}

func (f *TextField) Type() string {
	return "text"
}

func (f *TextField) Valid() (bool, string) {
	if f.Required && f.Value == "" {
		return false, "This field is required."
	}
	return true, ""
}

type PasswordField struct {
	BaseField
	Required bool
}

func (f *PasswordField) Type() string {
	return "password"
}

func (f *PasswordField) Valid() (bool, string) {
	if f.Required && f.Value == "" {
		return false, "This field is required."
	}
	return true, ""
}

type MatchingField struct {
	Field
	Matches Field
}

func (f *MatchingField) Valid() (bool, string) {
	if f.GetValue() != f.Matches.GetValue() {
		return false, "Fields do not match."
	}
	return true, ""
}

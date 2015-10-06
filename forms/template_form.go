package forms

import (
	"html/template"
)

type TemplateForm struct {
	Form
	Submitted bool
}

func (f TemplateForm) FieldHTML(name string) (template.HTML, error) {
	field, err := f.Get(name)
	if err != nil {
		return template.HTML(""), err
	}
	return template.HTML(field.HTML()), nil
}

func (f TemplateForm) FieldInvalidReason(name string) (string, error) {
	if f.Submitted {
		field, err := f.Get(name)
		if err != nil {
			return "", err
		}
		if valid, reason := field.Valid(); !valid {
			return reason, nil
		}
	}
	return "", nil
}

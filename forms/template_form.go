package forms

import (
	"html/template"
)

func NewTemplateForm() *TemplateForm {
	return &TemplateForm{New()}
}

type TemplateForm struct {
	*Form
}

func (f *TemplateForm) FieldHTML(name string) (template.HTML, error) {
	str, err := f.Form.FieldHTML(name)
	return template.HTML(str), err
}

func (f *TemplateForm) FieldInvalidReason(name string) (template.HTML, error) {
	str, err := f.Form.FieldInvalidReason(name)
	return template.HTML(str), err
}

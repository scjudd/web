package forms

import (
	"html/template"
)

type ErrFieldDoesNotExist struct {
	Name string
}

func (e *ErrFieldDoesNotExist) Error() string {
	return e.Name
}

type Form struct {
	Fields    map[string]Field
	Submitted bool
}

func New() *Form {
	return &Form{
		Fields:    make(map[string]Field),
		Submitted: false,
	}
}

func (f *Form) Valid() bool {
	for _, field := range f.Fields {
		if valid, _ := field.Valid(); !valid {
			return false
		}
	}
	return true
}

func (f *Form) FieldHTML(name string) (template.HTML, error) {
	if field, ok := f.Fields[name]; ok {
		if f.Submitted {
			if valid, _ := field.Valid(); !valid {
				field.AddClass("invalid")
			}
		}
		return FieldHTML(field, name), nil
	}
	return "", &ErrFieldDoesNotExist{name}
}

func (f *Form) FieldInvalidReason(name string) (template.HTML, error) {
	if field, ok := f.Fields[name]; ok {
		if f.Submitted {
			if valid, reason := field.Valid(); !valid {
				return template.HTML(reason), nil
			}
		}
		return template.HTML(""), nil
	}
	return template.HTML(""), &ErrFieldDoesNotExist{name}
}

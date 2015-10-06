package forms

import (
	"github.com/scjudd/web/forms/fields"
)

type ErrFieldDoesNotExist struct {
	Name string
}

func (err *ErrFieldDoesNotExist) Error() string {
	return "field '" + err.Name + "' does not exist"
}

type Form interface {
	Get(name string) (fields.Field, error)
	Field(fields.Field)
	Valid() bool
}

func New() Form {
	return make(form)
}

func Value(f Form, name string) string {
	field, err := f.Get(name)
	if err == nil {
		return field.GetValue()
	}
	return ""
}

type form map[string]fields.Field

func (f form) Get(name string) (fields.Field, error) {
	field, ok := f[name]
	if !ok {
		return nil, &ErrFieldDoesNotExist{name}
	}
	return field, nil
}

func (f form) Field(field fields.Field) {
	f[field.GetName()] = field
}

func (f form) Valid() bool {
	for _, field := range f {
		if valid, _ := field.Valid(); !valid {
			return false
		}
	}
	return true
}

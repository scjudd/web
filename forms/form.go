package forms

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

func (f *Form) FieldHTML(name string) (string, error) {
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

func (f *Form) FieldInvalidReason(name string) (string, error) {
	if field, ok := f.Fields[name]; ok {
		if f.Submitted {
			if valid, reason := field.Valid(); !valid {
				return reason, nil
			}
		}
		return "", nil
	}
	return "", &ErrFieldDoesNotExist{name}
}

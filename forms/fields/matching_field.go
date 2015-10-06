package fields

type MatchingField struct {
	Field
	Matching Field
	errors   []error
}

func (f *MatchingField) AddError(errs ...error) {
	for _, err := range errs {
		f.errors = append(f.errors, err)
	}
}

func (f *MatchingField) Valid() (bool, string) {
	if valid, reason := f.Field.Valid(); !valid {
		return false, reason
	}
	if f.Field.GetValue() != f.Matching.GetValue() {
		return false, "Fields must match."
	}
	if len(f.errors) > 0 {
		return false, f.errors[0].Error()
	}
	return true, ""
}

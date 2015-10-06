package fields

type Field interface {
	HTML() string
	GetName() string
	GetValue() string
	SetValue(string)
	AddError(...error)
	Valid() (bool, string)
}

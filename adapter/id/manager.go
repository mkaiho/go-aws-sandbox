package id

type IDGenerator interface {
	Generate() (string, error)
}

type IDValidator interface {
	Validate(v string) error
}

type IDManager interface {
	IDGenerator
	IDValidator
}

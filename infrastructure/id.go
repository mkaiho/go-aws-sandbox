package infrastructure

import (
	"github.com/mkaiho/go-aws-sandbox/adapter/id"
	"github.com/oklog/ulid/v2"
)

var _ id.IDGenerator = (*ULIDGenerator)(nil)
var _ id.IDValidator = (*ULIDValidator)(nil)
var _ id.IDManager = (*ULIDManager)(nil)

type ULIDGenerator struct{}

func (g *ULIDGenerator) Generate() (string, error) {
	return ulid.Make().String(), nil
}

type ULIDValidator struct{}

func (g *ULIDValidator) Validate(v string) error {
	_, err := ulid.Parse(v)
	if err != nil {
		return err
	}

	return err
}

type ULIDManager struct {
	generator *ULIDGenerator
	validator *ULIDValidator
}

func NewULIDManager() *ULIDManager {
	return &ULIDManager{
		generator: &ULIDGenerator{},
		validator: &ULIDValidator{},
	}
}

func (m *ULIDManager) Generate() (string, error) {
	return m.generator.Generate()
}

func (m *ULIDManager) Validate(v string) error {
	return m.validator.Validate(v)
}

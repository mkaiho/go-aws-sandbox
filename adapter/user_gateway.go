package adapter

import (
	"github.com/mkaiho/go-aws-sandbox/adapter/id"
	"github.com/mkaiho/go-aws-sandbox/entity"
	"github.com/mkaiho/go-aws-sandbox/usecase/port"
)

var _ (port.UserIDManager) = (*UserIDManager)(nil)

type UserIDManager struct {
	idManager id.IDManager
}

func NewUserIDManager(idManager id.IDManager) *UserIDManager {
	return &UserIDManager{
		idManager: idManager,
	}
}

func (m *UserIDManager) Generate() (entity.UserID, error) {
	id, err := m.idManager.Generate()
	if err != nil {
		return "", err
	}

	return entity.UserID(id), nil
}

func (m *UserIDManager) Parse(v string) (entity.UserID, error) {
	if err := m.idManager.Validate(v); err != nil {
		return "", err
	}

	return entity.UserID(v), nil
}

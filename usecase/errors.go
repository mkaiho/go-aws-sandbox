package usecase

import (
	"fmt"
)

// NotFoundEntityError
type NotFoundEntityError struct {
	targetName string
}

func NewNotFoundError(targetName string) *NotFoundEntityError {
	return &NotFoundEntityError{
		targetName: targetName,
	}
}

func (e *NotFoundEntityError) Error() string {
	return fmt.Sprintf("not found: %s", e.targetName)
}

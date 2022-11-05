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

func (e *NotFoundEntityError) TargetName() string {
	if e == nil {
		return ""
	}

	return e.targetName
}

func (e *NotFoundEntityError) Error() string {
	return fmt.Sprintf("not found: %s", e.targetName)
}

// DuplicateError
type DuplicateError struct {
	targetName string
}

func NewDuplicateError(targetName string) *DuplicateError {
	return &DuplicateError{
		targetName: targetName,
	}
}

func (e *DuplicateError) Error() string {
	return fmt.Sprintf("already exists: %s", e.targetName)
}

func (e *DuplicateError) TargetName() string {
	if e == nil {
		return ""
	}

	return e.targetName
}

package usecase

import "github.com/mkaiho/go-aws-sandbox/usecase/port"

var _ UserInteractor = (*userInteractorImpl)(nil)

func NewUserInteractor(userRepository port.UserRepository) *userInteractorImpl {
	return &userInteractorImpl{
		userRepository: userRepository,
	}
}

type UserInteractor interface{}

type userInteractorImpl struct {
	userRepository port.UserRepository
}

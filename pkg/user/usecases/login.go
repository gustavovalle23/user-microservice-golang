package usecases

import (
	"errors"
	"time"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token string
}

type LoginUseCase struct {
	userRepository domain.UserRepository
	tokenGenerator domain.TokenGenerator
}

func NewLoginUseCase(userRepository domain.UserRepository, tokenGenerator domain.TokenGenerator) LoginUseCase {
	return LoginUseCase{userRepository, tokenGenerator}
}

func (usecase LoginUseCase) Execute(input LoginInput) (LoginOutput, error) {
	user, err := usecase.userRepository.FindByEmail(input.Email)
	if err != nil {
		return LoginOutput{}, err
	}

	if err := user.ComparePassword(input.Password); err != nil {
		return LoginOutput{}, errors.New("Invalid password")
	}

	token, err := usecase.tokenGenerator.GenerateToken(user.ID, time.Now().Add(time.Hour*24*7))
	if err != nil {
		return LoginOutput{}, err
	}

	return LoginOutput{Token: token}, nil
}

package usecases

import (
	"errors"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

type CreateUserInput struct {
	Name       string
	Password   string
	Email      string
	DocumentNo string
	Address    domain.Address
	BirthDate  domain.Date
}

type CreateUserOutput struct {
	UserID string
}

type CreateUserUseCase struct {
	userRepo domain.UserRepository
}

func NewCreateUserUseCase(userRepo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInput) (CreateUserOutput, error) {
	existingUser, err := uc.userRepo.FindByDocumentNo(input.DocumentNo)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return CreateUserOutput{}, err
	}

	if existingUser != nil {
		return CreateUserOutput{}, domain.ErrUserAlreadyExists
	}

	user, err := domain.NewUser(nil, input.Name, input.Password, input.Email, input.DocumentNo, input.Address, input.BirthDate)
	if err != nil {
		return CreateUserOutput{}, err
	}

	err = uc.userRepo.Save(user)
	if err != nil {
		return CreateUserOutput{}, err
	}

	return CreateUserOutput{UserID: user.ID.Hex()}, nil
}

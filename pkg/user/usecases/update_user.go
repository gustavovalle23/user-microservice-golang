package usecases

import "github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"

type UpdateUserInput struct {
	Name      string
	Email     string
	Address   domain.Address
	BirthDate domain.Date
}

type UpdateUserOutput struct {
	UserID     string
	Name       string
	Email      string
	DocumentNo string
	Address    domain.Address
	BirthDate  domain.Date
}

type UpdateUserUseCase struct {
	userRepo domain.UserRepository
}

func NewUpdateUserUseCase(userRepo domain.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepo: userRepo}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInput) (UpdateUserOutput, error) {

}

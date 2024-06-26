package usecases

import (
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

type UpdateUserInput struct {
	UserID    int
	Name      string
	Email     string
	Address   domain.Address
	BirthDate domain.Date
}

type UpdateUserOutput struct {
	UserID     int
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
	originalUser, err := uc.userRepo.FindByID(input.UserID)
	if err != nil {
		return UpdateUserOutput{}, err
	}

	updatedUser, err := domain.NewUser(input.UserID, input.Name, originalUser.Password, input.Email, originalUser.DocumentNo, input.Address, input.BirthDate)
	if err != nil {
		return UpdateUserOutput{}, err
	}

	err = uc.userRepo.Update(updatedUser)
	if err != nil {
		return UpdateUserOutput{}, err
	}

	return UpdateUserOutput{
			UserID:     updatedUser.ID,
			Name:       updatedUser.Name,
			Email:      updatedUser.Email,
			DocumentNo: updatedUser.DocumentNo,
			Address:    updatedUser.Address,
			BirthDate:  updatedUser.BirthDate,
		},
		nil
}

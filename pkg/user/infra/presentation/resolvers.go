package presentation

import (
	"context"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/usecases"
)

type CreateUserPayload struct {
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

type mutationResolver struct {
	userRepo domain.UserRepository
}

func NewMutationResolver(userRepo domain.UserRepository) *mutationResolver {
	return &mutationResolver{userRepo: userRepo}
}

func (r *mutationResolver) CreateUser(ctx context.Context, input CreateUserPayload) (*CreateUserOutput, error) {
	createUserUseCase := usecases.NewCreateUserUseCase(r.userRepo)

	output, err := createUserUseCase.Execute(usecases.CreateUserInput{
		Name:       input.Name,
		Password:   input.Password,
		Email:      input.Email,
		DocumentNo: input.DocumentNo,
		Address:    input.Address,
		BirthDate:  input.BirthDate,
	})

	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		UserID: output.UserID,
	}, nil
}

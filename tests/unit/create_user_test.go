package tests

import (
	"errors"
	"testing"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/database"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/usecases"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUserUseCaseNewUserSuccess(t *testing.T) {
	userRepo := database.NewInMemoryUserRepository()

	input := usecases.CreateUserInput{
		Name:       "Tester Name",
		Password:   "password",
		Email:      "example@example.com",
		DocumentNo: "789012",
		Address: domain.Address{
			Street: "Tester Street",
			State:  "Tester State",
			City:   "Tester City",
		},
		BirthDate: domain.NewDate(1990, 1, 1),
	}

	useCase := usecases.NewCreateUserUseCase(userRepo)
	output, err := useCase.Execute(input)

	assert.NoError(t, err)
	assert.True(t, primitive.IsValidObjectID(output.UserID))

}

func TestCreateUserUseCaseExistingUserError(t *testing.T) {
	userRepo := database.NewInMemoryUserRepository()

	existingUser := &domain.User{
		Name:       "Test User",
		Password:   "password",
		Email:      "test@example.com",
		DocumentNo: "123456",
		Address: domain.Address{
			Street: "Random Street",
			State:  "Random State",
			City:   "Random City",
		},
		BirthDate: domain.NewDate(2000, 1, 1),
	}
	err := userRepo.Save(existingUser)
	assert.NoError(t, err)

	input := usecases.CreateUserInput{
		Name:       "Test User",
		Password:   "password",
		Email:      "test@example.com",
		DocumentNo: "123456",
		Address: domain.Address{
			Street: "Random Street",
			State:  "Random State",
			City:   "Random City",
		},
		BirthDate: domain.NewDate(2000, 1, 1),
	}

	expectedErr := domain.ErrUserAlreadyExists

	useCase := usecases.NewCreateUserUseCase(userRepo)
	output, err := useCase.Execute(input)

	assert.True(t, errors.Is(err, expectedErr))
	assert.Empty(t, output)
}

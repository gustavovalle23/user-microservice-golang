package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/usecases"
	"github.com/gustavovalle23/user-microservice-golang/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	email    = "email@email.com"
	name     = "Tester"
	password = "$2a$10$TgZJfh7eAtHsPTjIL8kGJ.rFv.tnYBlKjPbLZ2T45ZvTfM1tyQCZu" // hashed password for "secret123"
)

func TestLoginSuccessfulLoginReturnsToken(t *testing.T) {
	password := "secret123"
	expectedToken := "my-jwt-token"

	mockUserRepository := new(mocks.MockUserRepository)
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Password: password,
		Email:    email,
	}
	mockUserRepository.On("FindByEmail", email).Return(user, nil)

	mockTokenGenerator := new(mocks.MockTokenGenerator)
	expirationTime := time.Now().Add(time.Hour * 24)
	mockTokenGenerator.On("GenerateToken", user.ID.Hex(), expirationTime).Return(expectedToken, nil)

	loginUseCase := usecases.NewLoginUseCase(mockUserRepository, mockTokenGenerator)

	input := usecases.LoginInput{Email: email, Password: password}
	token, err := loginUseCase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedToken, token)
	mockUserRepository.AssertExpectations(t)
	mockTokenGenerator.AssertExpectations(t)
}

func TestLoginIncorrectPasswordReturnsError(t *testing.T) {
	password := "wrongpassword"

	mockUserRepository := new(mocks.MockUserRepository)
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Password: password,
		Email:    email,
	}
	mockUserRepository.On("FindByEmail", email).Return(user, nil)

	mockTokenGenerator := new(mocks.MockTokenGenerator)

	loginUseCase := usecases.NewLoginUseCase(mockUserRepository, mockTokenGenerator)

	// Act
	input := usecases.LoginInput{Email: email, Password: password}
	token, err := loginUseCase.Execute(input)

	// Assert
	assert.EqualError(t, err, "incorrect password")
	assert.Empty(t, token)
	mockUserRepository.AssertExpectations(t)
	mockTokenGenerator.AssertExpectations(t)
}

func TestLoginUserNotFoundReturnsError(t *testing.T) {
	password := "secret123"

	mockUserRepository := new(mocks.MockUserRepository)
	mockUserRepository.On("FindByEmail", email).Return(nil, errors.New("user not found"))

	mockTokenGenerator := new(mocks.MockTokenGenerator)

	loginUseCase := usecases.NewLoginUseCase(mockUserRepository, mockTokenGenerator)

	input := usecases.LoginInput{Email: email, Password: password}
	token, err := loginUseCase.Execute(input)

	// Assert
	assert.EqualError(t, err, "user not found")
	assert.Empty(t, token)
	mockUserRepository.AssertExpectations(t)
	mockTokenGenerator.AssertExpectations(t)
}

func TestLoginTokenGenerationFailsReturnsError(t *testing.T) {
	password := "secret123"
	mockUserRepository := new(mocks.MockUserRepository)

	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Password: password,
		Email:    email,
	}
	mockUserRepository.On("FindByEmail", email).Return(user, nil)

	mockTokenGenerator := new(mocks.MockTokenGenerator)
	mockTokenGenerator.On("GenerateToken", user.ID.Hex(), mock.Anything).Return("", errors.New("failed to generate token"))

	loginUseCase := usecases.NewLoginUseCase(mockUserRepository, mockTokenGenerator)

	// Act
	input := usecases.LoginInput{Email: email, Password: password}
	token, err := loginUseCase.Execute(input)

	// Assert
	assert.EqualError(t, err, "failed to generate token")
	assert.Empty(t, token)
	mockUserRepository.AssertExpectations(t)
	mockTokenGenerator.AssertExpectations(t)
}

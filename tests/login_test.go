package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByDocumentNo(documentNo string) (*domain.User, error) {
	args := m.Called(documentNo)
	user, _ := args.Get(0).(*domain.User)
	return user, args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	user, _ := args.Get(0).(*domain.User)
	return user, args.Error(1)
}

func (m *MockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

type MockTokenGenerator struct {
	mock.Mock
}

func (m *MockTokenGenerator) GenerateToken(userID string, expirationTime time.Time) (string, error) {
	args := m.Called(userID, expirationTime)
	return args.String(0), args.Error(1)
}

const email = "email@email.com"
const name = "Tester"
const password = "$2a$10$TgZJfh7eAtHsPTjIL8kGJ.rFv.tnYBlKjPbLZ2T45ZvTfM1tyQCZu" // hashed password for "secret123"

func TestLoginSuccessfulLoginReturnsToken(t *testing.T) {
	password := "secret123"
	expectedToken := "my-jwt-token"

	mockUserRepository := new(MockUserRepository)
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Password: password,
		Email:    email,
	}
	mockUserRepository.On("FindByEmail", email).Return(user, nil)

	mockTokenGenerator := new(MockTokenGenerator)
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

	mockUserRepository := new(MockUserRepository)
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Password: password,
		Email:    email,
	}
	mockUserRepository.On("FindByEmail", email).Return(user, nil)

	mockTokenGenerator := new(MockTokenGenerator)

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

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("FindByEmail", email).Return(nil, errors.New("user not found"))

	mockTokenGenerator := new(MockTokenGenerator)

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
	mockUserRepository := new(MockUserRepository)

	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Password: password,
		Email:    email,
	}
	mockUserRepository.On("FindByEmail", email).Return(user, nil)

	mockTokenGenerator := new(MockTokenGenerator)
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

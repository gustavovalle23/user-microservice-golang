package mocks

import (
	"time"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
	"github.com/stretchr/testify/mock"
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

func (m *MockUserRepository) FindByID(ID int) (*domain.User, error) {
	args := m.Called(ID)
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

func (m *MockTokenGenerator) GenerateToken(userID int, expirationTime time.Time) (string, error) {
	args := m.Called(userID, expirationTime)
	return args.String(0), args.Error(1)
}

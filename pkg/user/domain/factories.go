package domain

import (
	"time"
)

type UserOutput struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	IsStaff  bool   `json:"is_staff"`
}

func NewUser(userID int, name string, password string, email string, documentNo string, address Address, birthDate Date) (*User, error) {
	creationDate := time.Now().UTC()

	user := &User{
		ID:         userID,
		Name:       name,
		Password:   password,
		Email:      email,
		DocumentNo: documentNo,
		Address:    address,
		BirthDate:  birthDate,
		CreatedAt:  creationDate,
		UpdatedAt:  creationDate,
	}

	if err := user.EncryptPassword(); err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserOutput(userID int, name string, IsActive bool, IsStaff bool) (*UserOutput, error) {
	user := &UserOutput{
		ID:       userID,
		Name:     name,
		IsActive: IsActive,
		IsStaff:  IsStaff,
	}

	return user, nil
}

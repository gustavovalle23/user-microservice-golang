package domain

import (
	"time"
)

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

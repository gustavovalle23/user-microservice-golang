package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(name, password, email, documentNo string, address Address, birthDate Date) (*User, error) {
	creationDate := time.Now().UTC()

	user := &User{
		ID:         primitive.NewObjectID(),
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

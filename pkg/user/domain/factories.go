package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(userID *primitive.ObjectID, name string, password string, email string, documentNo string, address Address, birthDate Date) (*User, error) {
	creationDate := time.Now().UTC()

	if userID == nil {
		newID := primitive.NewObjectID()
		userID = &newID
	}

	user := &User{
		ID:         *userID,
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

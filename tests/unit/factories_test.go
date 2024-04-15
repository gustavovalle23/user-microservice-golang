package tests

import (
	"testing"
	"time"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewUser(t *testing.T) {
	name := "User Test"
	password := "password"
	email := "email@example.com"
	documentNo := "123456789"
	address := domain.Address{
		Street: "Street",
		City:   "City",
		State:  "State",
	}
	birthDate := domain.NewDate(1990, time.January, 1)

	t.Run("success", func(t *testing.T) {
		user, err := domain.NewUser(1, name, password, email, documentNo, address, birthDate)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, documentNo, user.DocumentNo)
		assert.Equal(t, address, user.Address)
		assert.Equal(t, birthDate, user.BirthDate)

		assert.NotEqual(t, primitive.NilObjectID, user.ID)

		err = user.ComparePassword(password)
		assert.NoError(t, err)

		assert.WithinDuration(t, time.Now().UTC(), user.CreatedAt, time.Second)
		assert.True(t, user.CreatedAt.Before(time.Now().UTC()))
		assert.True(t, user.UpdatedAt.Before(time.Now().UTC()))
		assert.WithinDuration(t, time.Now().UTC(), user.UpdatedAt, time.Second)
		assert.Nil(t, user.DeletedAt)
	})

}

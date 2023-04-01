package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Address struct {
	Street string
	State  string
	City   string
}

type User struct {
	ID         primitive.ObjectID
	Name       string
	BirthDate  Date
	Password   string
	Email      string
	DocumentNo string
	Address    Address
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func (d Date) Time() time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.UTC)
}

func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil && !u.DeletedAt.IsZero()
}

func (u *User) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

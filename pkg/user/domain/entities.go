package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Address struct {
	Street string
	State  string
	City   string
}

type User struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	BirthDate           Date
	Password            string
	Email               string
	DocumentNo          string
	Address             Address
	IsActive            bool `json:"is_active"`
	IsStaff             bool `json:"is_staff"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
	ResetToken          string
	ResetTokenExpiresAt time.Time
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

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

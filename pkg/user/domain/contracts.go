package domain

import "time"

type Mailer interface {
	SendPasswordResetEmail(to string) error
}

type TokenGenerator interface {
	GenerateToken(userID int, expirationTime time.Time) (string, error)
}

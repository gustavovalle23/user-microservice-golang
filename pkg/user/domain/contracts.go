package domain

import "time"

type Mailer interface {
	SendPasswordResetEmail(to string) error
}

type TokenGenerator interface {
	GenerateToken(userID string, expirationTime time.Time) (string, error)
}

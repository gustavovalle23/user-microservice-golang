package domain

type Mailer interface {
	SendPasswordResetEmail(to string) error
}

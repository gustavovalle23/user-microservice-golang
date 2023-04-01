package domain

type UserRepository interface {
	Save(user *User) error
	FindByDocumentNo(documentNo string) (*User, error)
}

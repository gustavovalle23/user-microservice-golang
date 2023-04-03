package domain

type UserRepository interface {
	Save(user *User) error
	FindByDocumentNo(documentNo string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
	Update(user *User) error
}

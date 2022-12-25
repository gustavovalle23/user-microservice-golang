package domain

func UserFactory(name string, email string, password string) User {
	return User{name: name, email: email, password: password, active: true, points: 0}
}

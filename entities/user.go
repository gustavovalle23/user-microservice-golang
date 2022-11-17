package entities

type User struct {
	name      string
	birthDate string
	points    int64
}

type IUser interface {
	GetName() string
	GetBirthDate() string
	GetPoins() int64
	IncreasePoints(points int64)
	ChangeName(name string)
	ChangeBirthDate(birthDate string)
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetBirthDate() string {
	return u.birthDate
}

func (u *User) GetPoins() int64 {
	return u.points
}

func (u *User) ChangeName(name string) {
	u.name = name
}

func (u *User) ChangeBirthDate(birthDate string) {
	u.birthDate = birthDate
}

func (u *User) IncreasePoints(points int64) {
	u.points = points
}

func UserFactory(name string, birthDate string, initialPoins int64) IUser {
	return &User{name: name, birthDate: birthDate, points: initialPoins}
}

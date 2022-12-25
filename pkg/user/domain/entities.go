package domain

type User struct {
	name      string
	email     string
	password  string
	active    bool
	points    int
	createdAt string
	updatedAt string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) IsActive() bool {
	return u.active
}

func (u *User) GetPoins() int {
	return u.points
}

func (u *User) GetCreatedAt() string {
	return u.createdAt
}

func (u *User) GetUpdatedAt() string {
	return u.updatedAt
}

func (u *User) IncreasePoints(points int) {
	u.points = points
}

func (u *User) Activate() {
	u.active = true
}

func (u *User) Deactivate() {
	u.active = false
}

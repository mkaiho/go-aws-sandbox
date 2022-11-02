package entity

type UserID string

func (i UserID) String() string {
	return string(i)
}

func NewUser(
	id UserID,
	name string,
	email string,
) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
	}
}

type User struct {
	id    UserID
	name  string
	email string
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

type Users []*User

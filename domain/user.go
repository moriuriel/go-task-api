package domain

import (
	"context"
	"time"
)

type (
	User struct {
		id        ID
		name      string
		email     string
		password  string
		createdAt time.Time
	}

	UserRepository interface {
		Create(User, context.Context) (User, error)
		FindById(string, context.Context) (User, error)
	}
)

func NewUser(id ID, name string, email string, password string, createdAt time.Time) User {
	return User{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		createdAt: createdAt,
	}
}

func (u User) ID() ID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

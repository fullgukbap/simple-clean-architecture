package domain

import (
	"context"
)

type User struct {
	ID          int
	Title       string
	Description string
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
	List(ctx context.Context, offset, limit int) ([]*User, error)
	Update(ctx context.Context, id int, user *User) (*User, error)
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
	List(ctx context.Context, offset, limit int) ([]*User, error)
	Update(ctx context.Context, id int, user *User) (*User, error)
	Delete(ctx context.Context, id int) error
}

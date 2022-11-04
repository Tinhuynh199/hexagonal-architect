package port

import (
	"context"

	. "hexrestapi1/internal/infrastructure/domain/user"
)


// Driven Actor -- Core -> MySQL DB
type UserRepository interface {
	GetAllUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *User) (int64, error)
	UpdateUser(ctx context.Context, user *User) (int64, error)
	DeleteUser(ctx context.Context, id string) (int64, error)
}
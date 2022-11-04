package service

import (
	"context"
	"database/sql"
	"log"

	. "hexrestapi1/internal/infrastructure/domain/user"
	. "hexrestapi1/internal/infrastructure/port/user"
)

type UserService interface {
	GetAllUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *User) (int64, error)
	UpdateUser(ctx context.Context, user *User) (int64, error)
	DeleteUser(ctx context.Context, id string) (int64, error)
}

type userService struct {
	DB         *sql.DB
	Repository UserRepository
}

func NewUserService(db *sql.DB, repos UserRepository) UserService {
	return &userService{DB: db, Repository: repos}
}

// GetAllUsers implements UserService
func (s *userService) GetAllUsers(ctx context.Context) (*[]User, error) {
	return s.Repository.GetAllUsers(ctx)
}

// GetUser implements UserService
func (s *userService) GetUser(ctx context.Context, id string) (*User, error) {
	return s.Repository.GetUser(ctx, id)
}

// CreateUser implements UserService
func (s *userService) CreateUser(ctx context.Context, user *User) (int64, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			log.Println(er)
			return -1, er
		}
		log.Println(er)
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

// UpdateUser implements UserService
func (s *userService) UpdateUser(ctx context.Context, user *User) (int64, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.Repository.UpdateUser(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			log.Println(err)
			return -1, err
		}
		log.Println(err)
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

// DeleteUser implements UserService
func (s *userService) DeleteUser(ctx context.Context, id string) (int64, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Fatal(err)
		return -1, nil
	}

	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.Repository.DeleteUser(ctx, id)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			log.Println(er)
			return -1, er
		}
		log.Println(err)
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

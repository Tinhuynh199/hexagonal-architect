package repository

import (
	"context"
	"database/sql"
	"log"

	. "hexrestapi1/internal/infrastructure/domain/user"
)

type UserAdapter struct {
	DB *sql.DB
}

func NewUserAdapter(db *sql.DB) *UserAdapter {
	return &UserAdapter{DB: db}
}

// GetAllUsers implements port.UserRepository
func (r *UserAdapter) GetAllUsers(ctx context.Context) (*[]User, error) {
	query := `select * from users`
	users := []User{}
	user := User{}
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		for rows.Next() {
			rows.Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.DateOfBirth)
			users = append(users, user)
		}
		return &users, nil
	}
}

// GetUser implements port.UserRepository
func (r *UserAdapter) GetUser(ctx context.Context, id string) (*User, error) {
	query := `
		select
			id, 
			username,
			email,
			phone,
			date_of_birth
		from users where id = ?`
	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Phone,
			&user.Email,
			&user.DateOfBirth)
		return &user, nil
	}
	return nil, nil
}

// CreateUser implements port.UserRepository
func (r *UserAdapter) CreateUser(ctx context.Context, user *User) (int64, error) {
	query := `
		insert into users (
			id,
			username,
			email,
			phone,
			date_of_birth)
		values (
			?,
			?,
			?, 
			?,
			?)`
	tx := GetTx(ctx)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx,
		user.ID,
		user.Username,
		user.Email,
		user.Phone,
		user.DateOfBirth)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

// UpdateUser implements port.UserRepository
func (r *UserAdapter) UpdateUser(ctx context.Context, user *User) (int64, error) {
	query := `
		update users 
		set
			username = ?,
			email = ?,
			phone = ?,
			date_of_birth = ?
		where id = ?`
	tx := GetTx(ctx)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, user.Username, user.Email, user.Phone, user.DateOfBirth, user.ID)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

// DeleteUser implements port.UserRepository
func (r *UserAdapter) DeleteUser(ctx context.Context, id string) (int64, error) {
	query := "delete from users where id = ?"
	tx := GetTx(ctx)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func GetTx(ctx context.Context) *sql.Tx {
	t := ctx.Value("tx")
	if t != nil {
		tx, ok := t.(*sql.Tx)
		if ok {
			return tx
		}
	}
	return nil
}

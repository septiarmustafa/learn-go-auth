package repository

import (
	"belajar-auth/domain"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type UserRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: goqu.New("default", con),
	}
}

// FindByID implements domain.UserRepository.
func (u *UserRepository) FindByID(ctx context.Context, id int64) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

// FindByUsername implements domain.UserRepository.
func (u *UserRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"username": username,
	})
	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

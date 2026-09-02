package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) SaveUser(ctx context.Context, user *domain.User) error {

	sqlStr := `INSERT INTO users(user_id, name, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repo.db.Exec(ctx, sqlStr, user.UserID, user.Name, user.UserName, user.Email, user.Password, time.Now())

	return err
}
func (repo *UserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	sqlStr := `SELECT id, name, username, email FROM users WHERE email = $1`

	err := repo.db.QueryRow(ctx, sqlStr, email).Scan(&user.UserID, &user.Name, &user.UserName, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("can't find user by email: %w", err)
	}

	return &user, nil
}

func (repo *UserRepo) FindByUserName(ctx context.Context, userName string) (*domain.User, error) {
	var user domain.User

	sqlStr := `SELECT user_id, name, username, email FROM users WHERE username = $1`

	err := repo.db.QueryRow(ctx, sqlStr, userName).Scan(&user.UserID, &user.Name, &user.UserName, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("can't find user by email: %w", err)
	}

	return &user, nil
}

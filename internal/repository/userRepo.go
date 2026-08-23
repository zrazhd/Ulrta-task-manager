package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) SaveUser(user *domain.User) error {
	projects, err := json.Marshal(user.Projects)
	if err != nil {
		return fmt.Errorf("cannot marshal user: %w", err)
	}

	sqlStr := `INSERT INTO users(id, name, username, email, password, projects) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = repo.db.Exec(context.Background(), sqlStr, user.UserID, user.Name, user.UserName, user.Email, user.Password, projects)

	return err
}
func (repo *UserRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	var rawProjects []byte

	sqlStr := `SELECT id, name, username, email, projects FROM users WHERE email = $1`

	err := repo.db.QueryRow(context.Background(), sqlStr, email).Scan(&user.UserID, &user.Name, &user.UserName, &user.Email, &rawProjects)
	if err != nil {
		return nil, fmt.Errorf("can't find user by email: %w", err)
	}
	err = json.Unmarshal(rawProjects, &user.Projects)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal users projects: %w", err)
	}

	return &user, nil
}

func (repo *UserRepo) FindByUserName(userName string) (*domain.User, error) {
	var user domain.User
	var rawProjects []byte

	sqlStr := `SELECT id, name, username, email, projects FROM users WHERE username = $1`

	err := repo.db.QueryRow(context.Background(), sqlStr, userName).Scan(&user.UserID, &user.Name, &user.UserName, &user.Email, &rawProjects)
	if err != nil {
		return nil, fmt.Errorf("can't find user by email: %w", err)
	}
	err = json.Unmarshal(rawProjects, &user.Projects)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal users projects: %w", err)
	}

	return &user, nil
}

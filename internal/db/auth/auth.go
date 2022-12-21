package auth

import (
	"context"

	"universe-auth/internal/errors"
	"universe-auth/internal/interfaces"
	"universe-auth/internal/models"
)

type AuthManager struct {
	interfaces.RowWorker
	interfaces.ExecWorker
}

func NewDbManager(rawWorker interface{}) *AuthManager {
	rowWorker := rawWorker.(interfaces.RowWorker)
	execWorker := rawWorker.(interfaces.ExecWorker)
	return &AuthManager{rowWorker, execWorker}
}

func (db *AuthManager) Registration(ctx context.Context, email, username, password string) (*models.User, error) {
	query := "insert into users (email,username, password) values ($1, $2, $3)"
	currentCount := 0
	db.QueryRow(ctx, "select count(1) from users where email =  $1", email).Scan(&currentCount)
	if currentCount > 0 {
		return nil, errors.AlreadyExists
	}
	_, err := db.Exec(ctx, query, email, username, password)
	if err != nil {
		return nil, err
	}

	user := models.User{}

	return &user, nil
}

func (db *AuthManager) Login(ctx context.Context, email, password string) (*models.User, error) {
	query := "Select username, email, password from users where email = $1"
	user := models.User{}

	db.QueryRow(ctx, query, email).Scan(&user.Username, &user.Email, &user.Password)
	return &user, nil
}

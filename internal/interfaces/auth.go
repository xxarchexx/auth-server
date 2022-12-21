package interfaces

import (
	"context"
	"universe-auth/internal/models"
)

type AuthWorker interface {
	Login(ctx context.Context, username, password string) (models.User, error)
}

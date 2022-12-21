package auth

import (
	"universe-auth/internal/db/auth"
	"universe-auth/internal/interfaces"
	"universe-auth/internal/jwt"
	"universe-auth/internal/models"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type AuthManager struct {
	db           DbAuthWorker
	tokenManager *jwt.TokenManager
}

type DbAuthWorker interface {
	Login(ctx context.Context, username, password string) (*models.User, error)
	Registration(ctx context.Context, email, username, password string) (*models.User, error)
}

func NewAuthManager(connector interfaces.DbConnector) (*AuthManager, error) {
	con, err := connector.GetDb()
	res := con.(*pgx.Conn)
	if err != nil {
		return nil, err
	}

	dbAuthManager := auth.NewDbManager(res)

	tmgr, err := jwt.New()
	if err != nil {
		return nil, err
	}

	return &AuthManager{db: dbAuthManager, tokenManager: tmgr}, nil
}

func (am *AuthManager) Validate(ctx context.Context, token string) (bool, error) {
	return am.tokenManager.Verify(token)
}

func (am *AuthManager) Registration(ctx context.Context, email, username, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return am.db.Registration(ctx, email, username, string(hashedPassword))
}

func (am *AuthManager) Login(ctx context.Context, email, password string) (*models.User, bool, error) {

	user, err := am.db.Login(context.Background(), email, password)
	if err != nil {
		return nil, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, false, nil
	}

	token, err := am.tokenManager.GenerateToken(user.Username, user.Email)
	if err != nil {
		return nil, false, err
	}

	user.Token = token
	return user, true, nil
}

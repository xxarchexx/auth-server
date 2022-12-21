package jwt

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func New() (*TokenManager, error) {
	data, err := os.ReadFile("./rsa_keys/private.key")
	if err != nil {
		return nil, err
	}

	pk, err := jwt.ParseRSAPrivateKeyFromPEM(data)
	if err != nil {
		return nil, err
	}

	return &TokenManager{
		privateKey: pk,
		publicKey:  &pk.PublicKey,
	}, nil
}

func (jm *TokenManager) GenerateToken(name, email string) (string, error) {
	return jm.sign(name, email)
}

func (jm *TokenManager) sign(name, email string) (string, error) {
	newToken := jwt.New(jwt.SigningMethodRS256)
	newToken.Claims = &UserInfoClaim{
		Name:             name,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token, err := newToken.SignedString(jm.privateKey)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (jm *TokenManager) Verify(token string) (bool, error) {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return jm.publicKey, nil
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

type UserInfoClaim struct {
	jwt.RegisteredClaims
	Name  string
	Email string
}

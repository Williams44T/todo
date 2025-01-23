package token_manager

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_SIGNING_METHOD  = jwt.SigningMethodHS256
	JWT_EXPIRATION_TIME = time.Duration(time.Minute * 5)
)

// assert that JWT_SIGNING_METHOD is of type SigningMethodHMAC
var _ jwt.SigningMethodHMAC = *JWT_SIGNING_METHOD

// TokenManager handles issuing and verifying JWTs.
type TokenManager struct {
	Secret []byte
}

// assert that TokenManager implements TokenManagerInterface
var _ TokenManagerInterface = &TokenManager{}

// NewTokenManager accepts a non-empty secret and returns a new instance of TokenManager.
func NewTokenManager(secret string) (*TokenManager, error) {
	if secret == "" {
		return nil, errors.New("secret cannot be empty")
	}
	return &TokenManager{Secret: []byte(secret)}, nil
}

// IssueToken creates a new jwt with the user ID as the subject that expires in 5 minutes.
// It signs the token with the token manager secret and returns it.
func (tm *TokenManager) IssueToken(userID string) (string, error) {
	if userID == "" {
		return "", errors.New("user ID cannot be empty")
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(JWT_EXPIRATION_TIME).Unix(),
	})

	signed, err := token.SignedString(tm.Secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %v", err)
	}

	return signed, nil
}

// verifySigningMethod is used to verify tokens by asserting the signing method is as expected.
func (tm *TokenManager) verifySigningMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return tm.Secret, nil
}

// VerifyToken asserts that the token is valid and returns the corresponding user id.
func (tm *TokenManager) VerifyToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, tm.verifySigningMethod)
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to extract claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("failed to extract user id from claims")
	}

	return userID, nil
}

package mock

import (
	"errors"
	"fmt"
	"todo/interfaces/token_manager"
)

// MockTokenManager mocks TockManager
type MockTokenManager struct {
	TokenMap       map[string]string
	count          int
	IssueTokenErr  error
	VerifyTokenErr error
}

// assert that MockTokenManager implements TokenManagerInterface
var _ token_manager.TokenManagerInterface = &MockTokenManager{}

// IssueToken creates new "token" by incrementing by 1 for each token (token_id_1, token_id_2, etc).
// It then adds that "token" to the TokenMap with the user ID as the value.
func (mtm *MockTokenManager) IssueToken(userID string) (string, error) {
	if mtm.TokenMap == nil {
		mtm.TokenMap = make(map[string]string)
	}
	mtm.count++
	token := fmt.Sprintf("token_id_%d", mtm.count)
	mtm.TokenMap[token] = userID
	return token, mtm.IssueTokenErr
}

// VerifyToken checks if the given token exists in the TokenMap and returns an error if not.
func (mtm *MockTokenManager) VerifyToken(token string) (string, error) {
	userID, ok := mtm.TokenMap["token"]
	if !ok {
		return "", errors.New("invalid token")
	}
	return userID, mtm.VerifyTokenErr
}

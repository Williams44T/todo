package token_manager

type TokenManagerInterface interface {
	IssueToken(userID string) (signedToken string, err error)
	VerifyToken(token string) (userID string, err error)
}

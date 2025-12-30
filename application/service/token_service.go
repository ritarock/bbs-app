package service

type TokenClaims struct {
	UserID int
}

type TokenService interface {
	Generate(userID int) (string, error)
	Verify(token string) (*TokenClaims, error)
}

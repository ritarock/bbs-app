package service

import (
	appservice "github.com/ritarock/bbs-app/application/service"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordService struct {
	cost int
}

func NewBcryptPasswordService() appservice.PasswordService {
	return &BcryptPasswordService{
		cost: bcrypt.DefaultCost,
	}
}

func (s *BcryptPasswordService) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *BcryptPasswordService) Verify(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

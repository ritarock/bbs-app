package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	appservice "github.com/ritarock/bbs-app/application/service"
)

type JWTTokenService struct {
	secretKey  []byte
	expiration time.Duration
}

func NewJWTTokenService(secretKey string, expiration time.Duration) appservice.TokenService {
	return &JWTTokenService{
		secretKey:  []byte(secretKey),
		expiration: expiration,
	}
}

func (s *JWTTokenService) Generate(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(s.expiration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *JWTTokenService) Verify(tokenString string) (*appservice.TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return nil, errors.New("invalid user_id claim")
		}
		return &appservice.TokenClaims{
			UserID: int(userID),
		}, nil
	}

	return nil, errors.New("invalid token")
}

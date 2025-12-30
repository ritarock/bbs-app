package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/service"
	"github.com/ritarock/bbs-app/application/usecase/auth"
	"github.com/ritarock/bbs-app/infra/handler/api"
)

type contextKey string

const requestContextKey contextKey = "request"

type AuthHandler struct {
	signUpUsecase         *auth.SignUpUsecase
	signInUsecase         *auth.SignInUsecase
	getCurrentUserUsecase *auth.GetCurrentUserUsecase
	tokenService          service.TokenService
}

func NewAuthHandler(
	signUpUsecase *auth.SignUpUsecase,
	signInUsecase *auth.SignInUsecase,
	getCurrentUserUsecase *auth.GetCurrentUserUsecase,
	tokenService service.TokenService,
) *AuthHandler {
	return &AuthHandler{
		signUpUsecase:         signUpUsecase,
		signInUsecase:         signInUsecase,
		getCurrentUserUsecase: getCurrentUserUsecase,
		tokenService:          tokenService,
	}
}

func (h *AuthHandler) AuthSignup(ctx context.Context, req *api.SignUpRequest) (*api.AuthResponse, error) {
	input := dto.SignUpInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	output, err := h.signUpUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.AuthResponse{
		User: api.User{
			ID:        int64(output.ID),
			Email:     output.Email,
			CreatedAt: output.CreatedAt,
		},
		Token: output.Token,
	}, nil
}

func (h *AuthHandler) AuthSignin(ctx context.Context, req *api.SignInRequest) (*api.AuthResponse, error) {
	input := dto.SignInInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	output, err := h.signInUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.AuthResponse{
		User: api.User{
			ID:        int64(output.ID),
			Email:     output.Email,
			CreatedAt: output.CreatedAt,
		},
		Token: output.Token,
	}, nil
}

func (h *AuthHandler) AuthMe(ctx context.Context) (*api.User, error) {
	req, ok := ctx.Value(requestContextKey).(*http.Request)
	if !ok {
		return nil, errors.New("unauthorized")
	}

	authHeader := req.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("unauthorized")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := h.tokenService.Verify(token)
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	input := dto.GetCurrentUserInput{
		UserID: claims.UserID,
	}

	output, err := h.getCurrentUserUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.User{
		ID:        int64(output.ID),
		Email:     output.Email,
		CreatedAt: output.CreatedAt,
	}, nil
}

func RequestContextKey() contextKey {
	return requestContextKey
}

package delivery

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/port"
)

type jwtCustomClaims struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type userHandler struct {
	userUsecase port.UserUsecase
}

func NewUserHandler(us port.UserUsecase) *userHandler {
	return &userHandler{
		userUsecase: us,
	}
}

func (u *userHandler) Signup(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	if err := user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := u.userUsecase.Create(ctx, &user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "success"})
}

func (u *userHandler) Login(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	loginUser, err := u.userUsecase.Find(ctx, user.Name, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	claims := &jwtCustomClaims{
		ID:   loginUser.ID,
		Name: loginUser.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func getIDByToken(c echo.Context) int {
	currentUser := c.Get("user").(*jwt.Token)
	claims := currentUser.Claims.(*jwtCustomClaims)
	return claims.ID
}

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := getIDByToken(c); id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
		}
		return next(c)
	}
}

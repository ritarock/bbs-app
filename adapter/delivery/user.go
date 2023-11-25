package delivery

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(us domain.UserUsecase) *userHandler {
	return &userHandler{
		userUsecase: us,
	}
}

func (u *userHandler) SignUp(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	ctx := c.Request().Context()

	if err := user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := u.userUsecase.SignUp(ctx, &user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, "success")
}

func (u *userHandler) Login(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	ctx := c.Request().Context()

	ok, loginUser := u.userUsecase.Login(ctx, &user)
	if !ok {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	if u.userUsecase.SetToken(ctx, loginUser.Id, t); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (u *userHandler) Session(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser := c.Get("user").(*jwt.Token)
		if !u.userUsecase.ValidateToken(c.Request().Context(), currentUser.Raw) {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		return next(c)
	}
}

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
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	ctx := c.Request().Context()

	if err := user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := u.userUsecase.Signup(ctx, &user); err != nil {
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

	loginUser, err := u.userUsecase.Login(ctx, user.Name, user.Password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		Name: loginUser.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	if u.userUsecase.UpdateToken(ctx, loginUser.ID, t); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (u *userHandler) Session(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser := c.Get("user").(*jwt.Token)
		if u.userUsecase.IsTokenAvailable(c.Request().Context(), currentUser.Raw) {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		return next(c)
	}
}

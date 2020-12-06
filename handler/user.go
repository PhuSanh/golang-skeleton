package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang-skeleton/dto"
	"golang-skeleton/dto/transform"
	"golang-skeleton/utils"
	"time"
)

// Login godoc
// @Summary Login
// @Description User login
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserLoginRequest true "Login"
// @Success 201 {object} dto.UserLoginResponse
// @Router /users/login [post]
func (h *Handler) Login(c echo.Context) error {
	req := new(dto.UserLoginRequest)
	if err := c.Bind(req); err != nil {
		return responseError(c, err)
	}

	user, err := h.userDomain.GetUserByUsername(req.Username)
	if err != nil {
		return responseError(c, err)
	}

	// Check password is match
	err = utils.CheckBCrypt(user.Password, req.Password)
	if err != nil {
		return responseError(c, err)
	}

	claims := &dto.JWTCustomClaims{
		Username: user.Username,
		ID:       user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(h.cfg.Jwt.ExpireInMinute)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := token.SignedString([]byte(h.cfg.Jwt.Secret))
	if err != nil {
		return responseError(c, err)
	}

	return responseSuccess(c, transform.ToUserLoginResponseDTO(user, encodedToken))
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "New User"
// @Success 201 {object} dto.CreateUserResponse
// @Router /users [post]
func (h *Handler) CreateUser(c echo.Context) error {
	req := new(dto.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return responseError(c, err)
	}

	user, err := h.userDomain.CreateUser(req)
	if err != nil {
		return responseError(c, err)
	}
	return responseSuccess(c, transform.ToCreateUserResponseDTO(user))
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 201 {object} dto.GetUserResponse
// @Router /users/{id} [get]
func (h *Handler) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	user, err := h.userDomain.GetUserByID(utils.IDStringToUint64(userID))
	if err != nil {
		return responseError(c, err)
	}
	return responseSuccess(c, transform.ToGetUserResponseDTO(user))
}

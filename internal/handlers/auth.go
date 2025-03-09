package handlers

import (
	"errors"
	"net/http"

	"social-network/internal/dto"
	"social-network/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHTTPHandlers struct {
	authService services.AuthService
}

func NewAuthHTTPHandlers(authService services.AuthService) AuthHandlers {
	return &AuthHTTPHandlers{authService: authService}
}

func (h *AuthHTTPHandlers) Login(c *gin.Context) {
	var userData dto.LoginUserDTO
	if err := c.BindJSON(&userData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenString, err := h.authService.LoginUser(userData)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) || errors.Is(err, services.ErrCannotFindUser) {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
			return
		} else if errors.Is(err, services.ErrCannotLoginUser) || errors.Is(err, services.ErrCannotSignToken) {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.IndentedJSON(http.StatusOK, tokenString)
}

func (h *AuthHTTPHandlers) Register(c *gin.Context) {
	var userData dto.UserCreateDTO
	if err := c.BindJSON(&userData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	createdUser, err := h.authService.CreateUser(userData)
	if err != nil {
		if errors.Is(err, services.ErrPasswordHashing) || errors.Is(err, services.ErrCannotCreateUser) {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		} else if errors.Is(err, services.ErrUserNotFound) {
			c.IndentedJSON(http.StatusNotFound, err.Error())
			return
		} else if errors.Is(err, services.ErrUserAlreadyExists) {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

	}
	c.IndentedJSON(http.StatusOK, createdUser)
}

func (h *AuthHTTPHandlers) GetMe(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := h.authService.ParseToken(tokenString)
	if err != nil {
		if errors.Is(err, services.ErrInvalidToken) {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
			return
		}
	}
	c.IndentedJSON(http.StatusOK, userId)
}

package user

import (
	"github.com/gin-gonic/gin"
	"go-server/utils"
	"net/http"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokenPayload := utils.TokenPayload{
		ID:       res.ID,
		Email:    res.Email,
		Username: res.Username,
	}
	tokenString, failed := utils.CreateToken(tokenPayload)

	if failed != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": failed.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]string{
		"token":    tokenString,
		"email":    res.Email,
		"username": res.Username,
	})
}

func (h *Handler) UserLogin(c *gin.Context) {
	var body LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.Service.UserLogin(c.Request.Context(), &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokenPayload := utils.TokenPayload{
		ID:       res.ID,
		Email:    res.Email,
		Username: res.Username,
	}
	tokenString, failed := utils.CreateToken(tokenPayload)

	if failed != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": failed.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token":    tokenString,
		"email":    res.Email,
		"username": res.Username,
	})
}

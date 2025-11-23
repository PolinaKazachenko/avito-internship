package users

import (
	"errors"
	"net/http"

	"avito-internship/internal/app/handlers/pkg"
	"avito-internship/internal/app/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *services.UserService
}

func NewHandler(userService *services.UserService) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) RegisterHandlers(g *gin.RouterGroup) {
	users := g.Group("users/")

	users.POST("setIsActive", h.SetIsActive)
	users.GET("getReview", h.GetReview)
}

func (h *Handler) SetIsActive(c *gin.Context) {
	var req SetIsActiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrInvalidRequestBody)
		return
	}
	user, err := h.userService.SetIsActive(c.Request.Context(), req.UserID, req.IsActive)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrInternalServerError)
	}
	c.JSON(http.StatusOK, convertIsActiveResponse(user))
}

func (h *Handler) GetReview(c *gin.Context) {
	teamName := c.Query("user_id")
	if teamName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrInvalidRequestBody)
		return
	}
	review, err := h.userService.GetReview(c.Request.Context(), teamName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrInternalServerError)
	}
	c.JSON(http.StatusOK, convertGetReviewResponse(review))
}

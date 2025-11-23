package teams

import (
	"errors"
	"fmt"
	"net/http"

	"avito-internship/internal/app/handlers/pkg"
	"avito-internship/internal/app/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	teamService *services.TeamService
}

func NewHandler(prService *services.TeamService) *Handler {
	return &Handler{teamService: prService}
}

func (h *Handler) RegisterHandlers(g *gin.RouterGroup) {
	team := g.Group("team/")

	team.POST("add", h.Add)
	team.GET("get", h.Get)
}

func (h *Handler) Add(c *gin.Context) {
	var req AddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrInvalidRequestBody)
		return
	}
	team, err := h.teamService.Add(c.Request.Context(), convertTeamAddRequest(&req))
	if err != nil {
		if errors.Is(err, services.ErrAlreadyExists) {
			msg := fmt.Sprintf("%s already exists", team.Name)
			c.AbortWithStatusJSON(http.StatusBadRequest, pkg.NewError(msg, "TEAM_EXISTS"))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, convertAddResponse(team))
}

func (h *Handler) Get(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrInvalidRequestBody)
		return
	}
	team, err := h.teamService.Get(c.Request.Context(), teamName)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, pkg.ErrNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrInternalServerError)
		return
	}
	c.JSON(http.StatusOK, convertTeam(team))
}

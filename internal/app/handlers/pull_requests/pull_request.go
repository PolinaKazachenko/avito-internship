package pull_requests

import (
	"errors"
	"net/http"

	"avito-internship/internal/app/handlers/pkg"
	"avito-internship/internal/app/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	prService *services.PullRequestService
}

func NewHandler(prService *services.PullRequestService) *Handler {
	return &Handler{prService: prService}
}

func (h *Handler) RegisterHandlers(g *gin.RouterGroup) {
	pullRequests := g.Group("pullRequest/")

	pullRequests.POST("create", h.Create)
	pullRequests.POST("merge", h.Merge)
	pullRequests.POST("reassign", h.Reassign)
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrInvalidRequestBody)
		return
	}
	pullRequest, err := h.prService.AddPullRequest(c.Request.Context(), convertCreateRequest(&req))
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrNotFound)
			return
		}
		if errors.Is(err, services.ErrAlreadyExists) {
			msg := "PR id already exists"
			c.AbortWithStatusJSON(http.StatusConflict, pkg.NewError(msg, "PR_EXISTS"))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrInternalServerError)
		return
	}
	c.JSON(http.StatusOK, convertAddResponse(pullRequest))
}

func (h *Handler) Merge(c *gin.Context) {
	var req MergeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrInvalidRequestBody)
		return
	}
	pullRequest, err := h.prService.MergePullRequest(c.Request.Context(), req.PullRequestID)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrInternalServerError)
		return
	}
	c.JSON(http.StatusOK, convertMergeResponse(pullRequest))
}

func (h *Handler) Reassign(c *gin.Context) {
	var req ReassignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrInvalidRequestBody)
		return
	}
	pullRequest, newReviewerID, err := h.prService.ReassignReviewer(
		c.Request.Context(),
		req.PullRequestID,
		req.OldReviewerID,
	)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrNotFound)
			return
		}
		if errors.Is(err, services.ErrPrAlreadyMerged) {
			c.AbortWithStatusJSON(http.StatusConflict,
				pkg.NewError(services.ErrPrAlreadyMerged.Error(), "PR_MERGED"),
			)
			return
		}
		if errors.Is(err, services.ErrReviewerNotAssignedToPR) {
			c.AbortWithStatusJSON(http.StatusConflict,
				pkg.NewError(services.ErrReviewerNotAssignedToPR.Error(), "NOT_ASSIGNED"))
			return
		}
		if errors.Is(err, services.ErrNoActiveCandidates) {
			c.AbortWithStatusJSON(http.StatusConflict,
				pkg.NewError(services.ErrNoActiveCandidates.Error(), "NO_CANDIDATE"))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrInternalServerError)
		return
	}
	c.JSON(http.StatusOK, convertReassignResponse(newReviewerID, pullRequest))
}

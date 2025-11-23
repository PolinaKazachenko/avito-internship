package users

type SetIsActiveRequest struct {
	UserID   string `json:"user_id" binding:"required,min=1"`
	IsActive bool   `json:"is_active" binding:"required"`
}

package request

type UserDataEmail struct {
	Email string `json:"email" binding:"required"`
}

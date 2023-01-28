package request

type Register struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Username        string `json:"username" binding:"required"`
}

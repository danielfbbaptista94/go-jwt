package requestdto

type SignupDTO struct {
	Email    string `json:"email" binding:"required,email" example:"test@email.com"`
	Password string `json:"password" binding:"required,min=6" example:"123456"`
}

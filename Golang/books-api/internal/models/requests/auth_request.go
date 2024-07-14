package requests

type SignupRequest struct {
	FullName        string `json:"full_name" validate:"required" example:"John Doe"`
	Username        string `json:"username" validate:"required" example:"johndoe"`
	Password        string `json:"password" validate:"required,gte=6,lte=20" example:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required,gte=6,lte=20,eqfield=Password" example:"password"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required" example:"johndoe"`
	Password string `json:"password" validate:"required,gte=6,lte=20" example:"password"`
}

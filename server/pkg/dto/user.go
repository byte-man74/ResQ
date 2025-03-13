package dto


type UserDTO struct {
	ID uint
	Email string
	FirstName string
	LastName string
}

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

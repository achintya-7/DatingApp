package dto

type CreateUserRequest struct {
	Email     string  `json:"email" binding:"required,email"`
	Name      string  `json:"name" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	Gender    string  `json:"gender" binding:"required"`
	Age       int8    `json:"age" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type CreateUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

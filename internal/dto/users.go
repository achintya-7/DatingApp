package dto

type CreateUserRequest struct {
	Email     string  `json:"email" binding:"required,email"`
	Name      string  `json:"name" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	Gender    string  `json:"gender" binding:"required"`
	Age       int8    `json:"age" binding:"required, gte=18, lte=100"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Age      int8   `json:"age"`
}

type DiscoverV2Request struct {
	Age struct {
		LessThan    *int32 `json:"less_than"`
		GreaterThan *int32 `json:"greater_than" binding:"gte=18"`
	} `json:"age"`
	Gender string `json:"gender,omitempty"`
}

type DiscoverV2Response struct {
	UserID   string  `json:"user_id"`
	Name     string  `json:"name"`
	Gender   string  `json:"gender"`
	Age      int32   `json:"age"`
	Distance float64 `json:"distance"`
}

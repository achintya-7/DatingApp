package v1

import (
	"log"

	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/logger"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (rh *RouteHandler) CreateUser(ctx *gin.Context) (*dto.CreateUserResponse, *dto.ErrorResponse) {
	var req dto.CreateUserRequest
	var resp dto.CreateUserResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "Error while binding request: ", err)
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Invalid request",
			HttpStatusCode: 400,
		}
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Error(ctx, "Error while hashing password: ", err)
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	user_id := uuid.New().String()

	args := db.CreateUserParams{
		UserID:    user_id,
		Email:     req.Email,
		Name:      req.Name,
		Gender:    req.Gender,
		Age:       int32(req.Age),
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Password:  password,
	}

	_, err = rh.store.CreateUser(ctx, args)
	if err != nil {
		log.Println("Error while creating user: ", err)
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	resp.ID = user_id
	resp.Email = req.Email
	resp.Name = req.Name

	return &resp, nil
}

func (rh *RouteHandler) DiscoverV1(ctx *gin.Context) (*[]db.DiscoverUsersV1Row, *dto.ErrorResponse) {
	authPayload := ctx.MustGet("authPayload").(*token.Payload)

	users, err := rh.store.DiscoverUsersV1(ctx, authPayload.UserId)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	if len(users) == 0 {
		return nil, &dto.ErrorResponse{
			Code:           404,
			Message:        "No users found",
			HttpStatusCode: 404,
		}
	}

	return &users, nil
}

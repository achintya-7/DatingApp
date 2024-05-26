package v2

import (
	"sort"

	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/internal/middleware"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
)

func (rh *RouteHandler) DiscoverV2(ctx *gin.Context) (*[]dto.DiscoverV2Response, *dto.ErrorResponse) {
	var req dto.DiscoverV2Request
	var resp []dto.DiscoverV2Response

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Invalid JSON body :" + err.Error(),
			HttpStatusCode: 400,
		}
	}

	authPayload := ctx.MustGet(middleware.AUTHORIZATION_PAYLOAD).(*token.Payload)

	discoverV2Args := db.DiscoverUsersV2Params{
		GreaterThanAge: utils.GetNullInt(req.Age.GreaterThan),
		LowerThanAge:   utils.GetNullInt(req.Age.LessThan),
		Gender:         utils.GetNullString(req.Gender),
		SwiperID:       authPayload.UserId,
	}

	users, err := rh.store.DiscoverUsersV2(ctx, discoverV2Args)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	currentUser, err := rh.store.GetUserByEmail(ctx, authPayload.Email)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	for _, user := range users {
		distance := utils.CalculateDistance(currentUser.Latitude, currentUser.Longitude, user.Latitude, user.Longitude, "K")

		respItem := dto.DiscoverV2Response{
			UserID:   user.UserID,
			Name:     user.Name,
			Gender:   user.Gender,
			Age:      user.Age,
			Distance: distance,
		}

		resp = append(resp, respItem)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Distance < resp[j].Distance
	})

	return &resp, nil
}

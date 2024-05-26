package v2

import (
	"github.com/achintya-7/dating-app/internal/dto"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
)

func (rh *RouteHandler) DiscoverV2(c *gin.Context) (*[]db.DiscoverUsersV2Row, *dto.ErrorResponse) {
	var req dto.DiscoverV2Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Invalid JSON body :" + err.Error(),
			HttpStatusCode: 400,
		}
	}

	discoverV2Args := db.DiscoverUsersV2Params{
		GreaterThanAge: utils.GetNullInt(req.Age.GreaterThan),
		LowerThanAge:   utils.GetNullInt(req.Age.LessThan),
		Gender:         utils.GetNullString(req.Gender),
	}

	users, err := rh.store.DiscoverUsersV2(c, discoverV2Args)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	return &users, nil
}

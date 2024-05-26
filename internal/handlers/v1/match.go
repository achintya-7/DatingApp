package v1

import (
	"net/http"

	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/logger"
	"github.com/gin-gonic/gin"
)

func (rh *RouteHandler) Match(ctx *gin.Context) (*gin.H, *dto.ErrorResponse) {
	var req dto.MatchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, &dto.ErrorResponse{
			Code:           http.StatusBadRequest,
			Message:        "Invalid request",
			HttpStatusCode: http.StatusBadRequest,
		}
	}

	

	logger.Info(ctx, "Match request received", req)

	resp := gin.H{
		"matched":  true,
		"match_id": "1234",
	}

	return &resp, nil
}

package db

import (
	"github.com/achintya-7/dating-app/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddMatchTx struct {
	User1Id string
	User2Id string
}

func (store *Store) AddMatchTx(ctx *gin.Context, arg AddMatchTx) (*CreateMatchParams, error) {
	var match CreateMatchParams
	
	err := store.ExecTx(ctx, func(q *Queries) error {

		// Create a swipe with YES for user1
		CreateSwipeArgs := CreateSwipeParams{
			SwipeID:   uuid.New().String(),
			SwiperID:  arg.User1Id,
			SwipeeID:  arg.User2Id,
			SwipeType: NullSwipesSwipeType{SwipesSwipeTypeYES, true},
		}

		_, err := q.CreateSwipe(ctx, CreateSwipeArgs)
		if err != nil {
			return err
		}

		// Check if user2 has swiped YES for user1
		CheckMatchArgs := CheckMatchParams{
			SwiperID:   arg.User2Id,
			SwipeeID:   arg.User1Id,
		}

		check, err := q.CheckMatch(ctx, CheckMatchArgs)
		if err != nil {
			return err
		}

		if !check {
			logger.Info(ctx, "No match found")
			return nil
		} else {
			logger.Info(ctx, "Match found")
		}

		// If YES, create a match
		CreateMatchArgs := CreateMatchParams{
			MatchID: uuid.New().String(),
			User1ID: arg.User1Id,
			User2ID: arg.User2Id,
		}

		match = CreateMatchArgs

		_, err = q.CreateMatch(ctx, CreateMatchArgs)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &match, nil
}

package worker

import (
	"context"
	"encoding/json"

	"github.com/achintya-7/dating-app/logger"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/worker"
	"github.com/hibiken/asynq"
)

func (processor *RedisTaskProcessor) SendMatchedEmailProcessor(ctx context.Context, task *asynq.Task) error {
	var payload worker.SendMatchEmailTaskPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	logger.Info(ctx, "processing send email task")

	// send email to user
	subject := "You have a new match!"
	body := worker.EmailBodyBuilder(payload.UserName, payload.MatchedUserName)
	body2 := worker.EmailBodyBuilder(payload.MatchedUserName, payload.UserName)

	// create a channel to handle errors
	errChan := make(chan error, 2)

	go func() {
		to := []string{payload.UserEmail}
		err := processor.mailer.SendEmail(subject, body, to, nil, nil, nil)
		errChan <- err
	}()

	go func() {
		to := []string{payload.MatchedUserEmail}
		err := processor.mailer.SendEmail(subject, body2, to, nil, nil, nil)
		errChan <- err
	}()

	// wait for both goroutines to finish and check for errors
	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}

func (processor *RedisTaskProcessor) CalculateUserAttractivenessProcessor(ctx context.Context, task *asynq.Task) error {
	var payload worker.CalculateUserAttractivenessTaskPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	logger.Info(ctx, "processing calculate user attractiveness task")

	// calculate user attractiveness
	userRanking, err := processor.store.GetRankingByUserId(ctx, payload.Userid)
	if err != nil {
		return err
	}

	// ? Update user ranking
	//  For now I am following a simple logic to update the ranking.
	//  The attractiveness of the user is calculated based on the response with the following formula:
	//  attractiveness = likes / (likes + dislikes)

	likes := userRanking.LikeCount
	dislikes := userRanking.DislikeCount
	var attractiveness float64

	if payload.Response == "YES" {
		likes++
	} else {
		dislikes++
	}

	if likes+dislikes == 0 {
		attractiveness = 0
	} else {
		attractiveness = float64(likes) / float64(likes+dislikes)
	}

	// update the user ranking
	updateRankingArgs := db.UpdateRankingParams{
		LikeCount:           likes,
		DislikeCount:        dislikes,
		AttractivenessScore: attractiveness,
		UserID:              payload.Userid,
	}

	_, err = processor.store.UpdateRanking(ctx, updateRankingArgs)
	if err != nil {
		return err
	}

	return nil
}
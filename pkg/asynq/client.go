package asynqclient

import (
	"github.com/achintya-7/dating-app/config"
	"github.com/hibiken/asynq"
)

func NewClient() *asynq.Client {
	return asynq.NewClient(
		asynq.RedisClientOpt{Addr: config.Values.RedisUrl},
	)
}

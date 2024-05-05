package queue

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/go-redis/redis/v8"
	"github.com/jaquelineabreu/to-do/models"
)

type QueueName string

const (
	ListToDO   QueueName = "ListToDO"
	StatusList QueueName = "StatusList"
)

func PublishList(ctx context.Context, rd *redis.Client, queue QueueName, list models.ToDoList) error {

	slog.Info("Publicando lista no Redis", slog.String("queue", string(queue)))

	data, err := json.Marshal(list)
	if err != nil {
		slog.Error("Erro ao converter lista para JSON", slog.String("erro", err.Error()))
		return err
	}

	if err := rd.RPush(ctx, string(queue), data).Err(); err != nil {
		slog.Error("Erro ao adicionar a lista ao Redis", slog.String("queue", string(queue)), slog.String("erro", err.Error()))
		return err
	}

	slog.Info("Lista publicada com sucesso", slog.String("queue", string(queue)))
	return nil
}

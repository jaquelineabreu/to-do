package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jaquelineabreu/to-do/models"
	"github.com/jaquelineabreu/to-do/queue"
	"github.com/jaquelineabreu/to-do/server"
)

func main() {

	redis, err := server.ConnectRedis()
	if err != nil {
		panic(err)
	}

	list := models.ToDoList{
		Id:         1,
		Title:      "Lista de compras",
		List:       []string{"Arroz", "Feijão", "Batata", "Sabonete", "Creme dental"},
		StatusList: models.EmAndamento,
	}

	fmt.Println(list)

	ctx := context.Background()
	err = queue.PublishList(ctx, redis, queue.ListToDO, list)
	if err != nil {
		log.Fatalf("Erro ao adicionar a lista ao Redis: %v", err)
	}

}

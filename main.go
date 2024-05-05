package main

import (
	"fmt"

	"github.com/jaquelineabreu/to-do/models"
)

func main() {

	list := models.ToDoList{
		Title:      "Lista de compras",
		List:       []string{"Arroz", "Feijão", "Batata", "Sabonete", "Creme dental"},
		StatusList: models.EmAndamento,
	}

	fmt.Println(list)

}

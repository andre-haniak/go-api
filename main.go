package main

import (
	"fmt"
	"github.com/andre-haniak/go-api/models"
	"github.com/andre-haniak/go-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Primeiro nome", History: "Historia 1"},
		{Name: "Segundo nome", History: "Historia 2"},
	}

	fmt.Println("Iniciando Servidor!!!")
	routes.HandleRequest()
}

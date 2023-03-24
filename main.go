package main

import (
	"fmt"
	"go-api/models"
	"go-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Primeiro nome", History: "Historia 1"},
		{Name: "Segundo nome", History: "Historia 2"},
	}

	fmt.Println("Iniciando Servidor!!!")
	routes.HandleRequest()
}

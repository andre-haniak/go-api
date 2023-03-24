package main

import (
	"fmt"
	"go-api/routes"
)

func main() {
	fmt.Println("Iniciando Servidor!!!")
	routes.HandleRequest()
}

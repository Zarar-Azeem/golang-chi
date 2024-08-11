package main

import (
	"fmt"

	"github.com/Zarar-Azeem/golang-chi/pkg/config"
	"github.com/Zarar-Azeem/golang-chi/pkg/models"
)

func main() {
	port := ":3000"
	server := config.NewApiServer(port)
	models.Init()

	fmt.Printf("It is running at localhost%v", port)
	server.Run()
}

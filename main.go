package main

import (
	"github.com/Henriquemcc/api-go-gin/database"
	"github.com/Henriquemcc/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

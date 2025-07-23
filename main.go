package main

import (
	"github.com/Henriquemcc/api-go-gin/database"
	"github.com/Henriquemcc/api-go-gin/models"
	"github.com/Henriquemcc/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "João da Silva", CPF: "893.754.630-28", RG: "44.145.525-6"},
		{Nome: "José da Silva", CPF: "547.444.120-76", RG: "36.602.077-8"},
	}
	routes.HandleRequests()
}

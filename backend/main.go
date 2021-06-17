package main

import (
	config "homefill/backend/config"
	"homefill/backend/db"
	"homefill/backend/routes"
)

func main() {
	config.LoadConfig()

	pgRepository := &db.PostgreSQL{}

	db.DB = db.DataBaseRepo{
		Repo: pgRepository,
	}
	db.DB.Repo.ConnectTODB()
	db.DB.Repo.RunDBScripts()

	routes.StartServer()
}

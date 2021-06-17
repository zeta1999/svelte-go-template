package db

import (
	"database/sql"
	"homefill/backend/model"
)

type DatabaseInterface interface {
	ConnectTODB()
	RunDBScripts()
	GetUserFromId(string) (*model.User, error)
	InsertUser(*model.User) error
}

type PostgreSQL struct {
	db *sql.DB
}

type DataBaseRepo struct {
	Repo DatabaseInterface
}

var DB DataBaseRepo

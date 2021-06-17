package service

import (
	"encoding/json"
	"homefill/backend/db"
)

func GetUserById(id string) ([]byte, error) {
	user, err := db.DB.Repo.GetUserFromId(id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}

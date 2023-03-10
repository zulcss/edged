package db

import (
        database "github.com/zulcss/edged/shared/db"
        "github.com/zulcss/edged/shared/model"
)

func CreateServer(s *model.Server) (err error) {
	if err = database.DB.Create(s).Error; err != nil {
		return err
	}
	return nil
}

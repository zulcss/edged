package db

import (
	database "github.com/zulcss/edged/shared/db"
	"github.com/zulcss/edged/shared/model"
)

func CreateSite(s *model.Site) (err error) {
	if err = database.DB.Create(s).Error; err != nil {
		return err
	}
	return nil
}

func ListSites(s *[]model.Site) (err error) {
	if err = database.DB.Find(&s).Error; err != nil {
		return err
	}
	return nil
}

func GetSite(s *model.Site, name string) (err error) {
	if err := database.DB.Where("name=?", name).First(&s).Error; err != nil {
		return err
	}
	return nil
}

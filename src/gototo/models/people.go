package models

import (
	"github.com/jinzhu/gorm"
	"gototo/db"
	"log"
)

type People struct {
	gorm.Model;

	UserId int
	Name string  `gorm:"size:255"`
	Address string `gorm:"size:255"`
}

func GetPeople() ([]People, error){
	var people []People
	err := db.GetDb().Table("peoples").Scan(&people).Error
	if err != nil {
		log.Println("errr", err)
		return people, err
	}

	return people, nil
}

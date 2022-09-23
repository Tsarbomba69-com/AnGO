package models

import (
	"gorm.io/gorm"
)

type Province struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);unique" json:"nome"` // unique, indexes, keys, etc cannot be to long
}

func IntialData(db *gorm.DB) {
	var provinces = []Province{
		{Name: "Luanda"},
		{Name: "Cabinda"},
		{Name: "Cunene"},
		{Name: "Benguela"},
	}
	db.Create(&provinces)
}

func GetProvince() Province {
	var province Province
	return province
}

func GetProvinces() []Province {
	var provinces []Province
	return provinces
}

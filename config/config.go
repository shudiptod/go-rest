package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() (*gorm.DB,error){
	db,err:= gorm.Open("postgres","user=shudipto password=shudipto dbname=pcbuilderdb sslmode=disable")
	if err != nil {
		return nil, err
	}
	DB= db
	return db,nil
}
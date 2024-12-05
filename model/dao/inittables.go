package dao

import (
	"log"
	"our_blog/db"
)

func InitTables() error {
	err := db.DB.AutoMigrate(&Article{}, &User{})
	if err != nil {
		log.Println("InitTables failed, err:", err)
		return err
	}
	return err
}

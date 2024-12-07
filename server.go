package main

import (
	"log"
	"our_blog/db"
	"our_blog/model/dao"
	"our_blog/route"
)

func main() {

	if err := InitDB(); err != nil {
		log.Println("init db failed, err:", err)
	}

	r := route.SetUpRouter()
	r.Run()

}

func InitDB() (err error) {
	if err = db.InitMysql(); err != nil {
		return err
	}
	if err = db.InitRedis(); err != nil {
		return err
	}
	if err = dao.InitTables(); err != nil {
		return err
	}
	return nil
}

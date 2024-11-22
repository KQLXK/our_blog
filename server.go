package main

import (
	"our_blog/repository"
	"our_blog/route"
)

func main() {

	if err := InitDB(); err != nil {
		panic(err)
	}

	r := route.SetUpRouter()
	r.Run()

}

func InitDB() (err error) {
	if err = repository.InitMysql(); err != nil {
		return err
	}
	return nil
}

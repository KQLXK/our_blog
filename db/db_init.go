package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
	Ctx = context.Background()
)

func InitMysql() (err error) {
	dsn := "root:yuchao123@tcp(127.0.0.1:3306)/ourblog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func InitRedis() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("could not connect to redis: %v", err)
		return err
	}
	fmt.Println("connect success")
	return nil
}

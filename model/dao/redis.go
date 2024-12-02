package dao

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"our_blog/db"
	"sync"
)

var tokenDao *RedisDao
var tokenOnce sync.Once

type RedisDao struct {
	Rdb *redis.Client
	Ctx context.Context
}

func NewTokenDao() *RedisDao {
	tokenOnce.Do(func() {
		tokenDao = &RedisDao{
			Rdb: db.Rdb,
			Ctx: context.Background(),
		}
	})
	return tokenDao
}

func (r *RedisDao) SetKey(key string, value string) error {

	if r.Rdb == nil {
		log.Println("Rdb is nil, please check Redis connection")
		return errors.New("Rdb is nil")
	}

	err := r.Rdb.Set(r.Ctx, key, value, 0).Err()
	if err != nil {
		log.Println("redis set key failed, err:", err)
		return err
	}

	log.Println("set key sucessfully, key:", key, "value:", value)
	return nil
}

func (r *RedisDao) GetKey(key string) (string, error) {
	val, err := r.Rdb.Get(r.Ctx, key).Result()
	if err != nil {
		log.Println("redis get key failed, err:", err)
		return "", err
	}
	return val, nil
}

func (r *RedisDao) DelKey(key string) error {
	err := r.Rdb.Del(r.Ctx, key).Err()
	if err != nil {
		log.Println("redis Del key failed, err:", err)
		return err
	}
	return nil
}

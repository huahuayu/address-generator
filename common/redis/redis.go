package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/huahuayu/address-generator/common/config"
	"time"
)

var (
	Client *redis.Client
)

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.App.Redis.Host,
		Password: config.App.Redis.Pass,
		DB:       config.App.Redis.Db,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func ObtainLock(key string, expiration time.Duration) error {
	val, err := Client.SetNX(key, 1, expiration).Result()
	if err != nil {
	}
	if !val {
		return errors.New("lock is been taken")
	}
	return nil
}

func ReleaseLock(key string) {
	Client.Del(key)
}

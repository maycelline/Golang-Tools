package goredis

import (
	"GoTools/model"
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

func GetUser(user model.User) {
	converted, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	err = client.Set(ctx, "user", converted, 0).Err()
	if err != nil {
		log.Println(err)
		return
	}

	value, err := client.Get(ctx, "user").Result()
	if err != nil {
		log.Println(err)
		return
	}

	_ = json.Unmarshal([]byte(value), &user)
	log.Println(user)
}

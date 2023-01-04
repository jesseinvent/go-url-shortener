package database

import (
	"fmt"
	"errors"
	"time"
	"context"
	"os"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background();

var rdb = redis.NewClient(&redis.Options{
	Addr: os.Getenv("DB_URL"),
	Password: "",
});



// func CreateClient(dbNo int) *redis.Client {

// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: os.Getenv("DB_ADDR"),
// 		Password: "",
// 		DB: dbNo,
// 	});

// 	return rdb;
// }

func getFullKey(key string) string {
	return "urlShortener:" + key;
}

func Set(key string, value string, ttl time.Duration) (bool, error) {

	err := rdb.Set(Ctx, getFullKey(key), value, ttl).Err();

	if err != nil {
		fmt.Println(err)
		return false, err;
	}

	return true, nil;
}

func Get(key string) (string, error) {

	value, err := rdb.Get(Ctx, getFullKey(key)).Result();

	if err == redis.Nil {
		return "", errors.New("no result found");
	} 

	return value, nil;
}

func TTL(key string) (time.Duration, error) {

	value, err := rdb.TTL(Ctx, getFullKey(key)).Result();

	if err == redis.Nil {
		return 0, errors.New("no result found");
	}

	return value, nil;
}


func IncrementKeyValue(key string) (int, error) {
	value, err := rdb.Incr(Ctx, getFullKey(key)).Result();

	if err == redis.Nil {
		return 0, errors.New("no result found");
	}

	return int(value), nil;
}

func DecrementKeyValue(key string) (int, error) {
	value, err := rdb.Decr(Ctx, getFullKey(key)).Result();

	if err == redis.Nil {
		return 0, errors.New("no result found");
	}

	return int(value), nil;

}
package database

import (
    "fmt"
    "context"

    "kolkata/internal/config"

    "github.com/go-redis/redis/v9"
)

type redisDB struct {
    client    *redis.Client
}

func NewRedis(conf config.RedisConfig) (Database, error) {
    addr := fmt.Sprintf("%v:%v", conf.Addr, conf.Port)
    client := redis.NewClient(&redis.Options{
        Addr: addr,
        Password: "",
        DB: 0,
    })

    ctx := context.Background()

    err := client.Ping(ctx).Err()
    if err != nil {
        return nil, &ErrCreateDatabase{}
    }

    return &redisDB{client}, nil
}

func (rdb *redisDB) Set(key string, val string) (string, error) {
    ctx := context.Background()

    err := rdb.client.Set(ctx, key, val, 0).Err()
    if err == redis.Nil {
        return "", &ErrOperation{"set"}
    } else if err != nil {
        return "", &ErrConnectDatabase{}
    }

    return key, nil
}

func (rdb *redisDB) Get(key string) (string, error) {
    ctx := context.Background()

    val, err := rdb.client.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", &ErrOperation{"get"}
    } else if err != nil {
        return "", &ErrConnectDatabase{}
    }

    return val, nil
}

func (rdb *redisDB) Delete(key string) (string, error) {
    ctx := context.Background()

    err := rdb.client.Del(ctx, key).Err()
    if err == redis.Nil {
        return "", &ErrOperation{"delete"}
    } else if err != nil {
        return "", &ErrConnectDatabase{}
    }

    return key, nil
}


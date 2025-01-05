package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // 没有密码，默认值
	DB:       0,  // 默认DB 0
})
var ctx = context.Background()

func Get(key string) string {
	val, err := client.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key不存在")
		return ""
	case err != nil:
		fmt.Println("错误", err)
		return err.Error()
	case val == "":
		fmt.Println("值是空字符串")
		return val
	}
	return val
}
func Set(key string, val string) bool {
	err := client.Set(ctx, key, val, 0)
	if err != nil {
		fmt.Println(err.Err())
		return false
	}
	return true
}
func LLen(key string) int64 {
	return 0
}

/*
// 设置一个键值对
err = rdb.Set(ctx, "key", "value", 0).Err()
if err != nil {
log.Fatalf("Failed to set key: %v", err)
}

// 获取键的值
val, err := rdb.Get(ctx, "key").Result()
if err != nil {
log.Fatalf("Failed to get key: %v", err)
}
fmt.Println("Key:", val)

// 使用哈希
err = rdb.HSet(ctx, "hashKey", "field", "value").Err()
if err != nil {
log.Fatalf("Failed to set hash field: %v", err)
}

val, err = rdb.HGet(ctx, "hashKey", "field").Result()
if err != nil && err != redis.Nil {
log.Fatalf("Failed to get hash field: %v", err)
} else if err == redis.Nil {
val = "nil"
}
fmt.Println("Hash field:", val)

// 使用列表
err = rdb.LPush(ctx, "listKey", "value1").Err()
if err != nil {
log.Fatalf("Failed to push to list: %v", err)
}

err = rdb.LPush(ctx, "listKey", "value2").Err()
if err != nil {
log.Fatalf("Failed to push to list: %v", err)
}

vals, err := rdb.LRange(ctx, "listKey", 0, -1).Result()
if err != nil {
log.Fatalf("Failed to get list range: %v", err)
}
fmt.Println("List:", vals)

// 使用带有过期时间的键值对
err = rdb.SetEX(ctx, "expireKey", "expireValue", 10*time.Second).Err()
if err != nil {
log.Fatalf("Failed to set expirable key: %v", err)
}

// 等待一段时间后检查过期键
time.Sleep(11 * time.Second)
val, err = rdb.Get(ctx, "expireKey").Result()
if err == redis.Nil {
fmt.Println("Expirable key has expired")
} else if err != nil {
log.Fatalf("Failed to get expirable key: %v", err)
} else {
fmt.Println("Expirable key:", val)
}

// 关闭 Redis 客户端连接
rdb.Close()*/

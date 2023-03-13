package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "172.16.15.98:6379",
	Password: "qwe123-=",
	DB:       0, // use default DB
})

func TestRedisSet(t *testing.T) {
	err := rdb.Set(ctx, "key", "123", 0).Err()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisGet(t *testing.T) {
	result, err := rdb.Get(ctx, "login_tokens").Result()
	if err != nil {
		fmt.Println("this is err", err)
		t.Fatal(err)
	}
	t.Log(result)
}

func TestTimeFormat(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(1676475453147, 0).Format(timeLayout)
	fmt.Println(timeStr)
}

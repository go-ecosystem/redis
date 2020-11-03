package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

var client *redis.Client

// GetClient get redis client
func GetClient() *redis.Client {
	return client
}

// Close closes the client, releasing any open resources
func Close() error {
	return client.Close()
}

// SetUp set up redis connection
func SetUp(host, port, pwd string) {
	ops := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pwd,
		DB:       0, // use default DB
	}
	SetUpWithOps(ops)
}

// SetUpWithOps set up redis connection with Options
func SetUpWithOps(ops *redis.Options) {
	client = redis.NewClient(ops)
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed to connect to Redis", pong, err)
		panic(err)
	}
	log.Println("Successfully connect to Redis")
}

// MockRedis mock redis for unittests
func MockRedis() {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

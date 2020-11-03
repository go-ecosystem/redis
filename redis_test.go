package redis

import (
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/go-redis/redis/v8"
)

func TestGetClient(t *testing.T) {
	ops := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),
		Password: "123456",
		DB:       0,
	}

	client = redis.NewClient(ops)

	if got := GetClient(); !reflect.DeepEqual(got, client) {
		t.Errorf("GetClient() = %v, want %v", got, client)
	}

	client = nil

	if got := GetClient(); got != nil {
		t.Errorf("GetClient() = %v, want %v", got, nil)
	}
}

func TestClose(t *testing.T) {
	MockRedis()

	monkey.Patch(client.Close, func() error { return nil })

	var want error = nil
	got := Close()

	if got != nil {
		t.Errorf("Close() = %v, want %v", got, want)
	}

	monkey.Unpatch(client.Close)

	monkey.Patch(client.Close, func() error { return fmt.Errorf("error") })

	got = Close()

	if got == nil {
		t.Errorf("Close() error")
	}
}

func TestSetUp(t *testing.T) {
	var tmp *redis.Options

	monkey.Patch(SetUpWithOps, func(ops *redis.Options) { tmp = ops })

	const (
		host = "127.0.0.1"
		port = "11234"
		pwd  = "dasdasd"
	)

	SetUp(host, port, pwd)

	want := fmt.Sprintf("%s:%s", host, port)
	if tmp.Addr != want {
		t.Errorf("SetUp() address %v, want %v", tmp.Addr, want)
	}
}

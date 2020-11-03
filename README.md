# redis

![Go](https://github.com/go-ecosystem/redis/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/go-ecosystem/redis/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ecosystem/redis)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ecosystem/redis)](https://goreportcard.com/report/github.com/go-ecosystem/redis)
[![Release](https://img.shields.io/github/release/go-ecosystem/redis.svg)](https://github.com/go-ecosystem/redis/releases)

> Golang redis package based on [go-redis](https://github.com/go-redis/redis) and [miniredis](https://github.com/alicebob/miniredis).

## Installation

```shell
go get github.com/go-ecosystem/redis
```

## Usage

See more at [go-redis](https://github.com/go-redis/redis).

```go
package example_test

import (
	"context"
	"fmt"

	"github.com/go-ecosystem/redis"
	goredis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleGetClient() {
	redis.SetUp("127.0.0.1", "6379", "")
	defer redis.Close()

	err := redis.GetClient().Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redis.GetClient().Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := redis.GetClient().Get(ctx, "key2").Result()
	if err == goredis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
```
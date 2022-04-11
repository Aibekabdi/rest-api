package repository

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.TODO()

type CounterRedis struct {
	client *redis.Client
}

func NewCounterRedis(client *redis.Client) *CounterRedis {
	return &CounterRedis{client: client}
}

func (c *CounterRedis) Add(num string) error {
	val, err := c.client.Get(ctx, "counter").Result()
	if err != nil || val == "" {
		c.client.Set(ctx, "counter", 0, 0)
	}
	numValue, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	numAdder, err := strconv.Atoi(num)
	if err != nil {
		return err
	}
	res := numValue + numAdder
	c.client.Set(ctx, "counter", res, 0)
	return nil
}

func (c *CounterRedis) Sub(num string) error {
	val, err := c.client.Get(ctx, "counter").Result()
	if err != nil || val == "" {
		c.client.Set(ctx, "counter", 0, 0)
	}
	numValue, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	numAdder, err := strconv.Atoi(num)
	if err != nil {
		return err
	}
	res := numValue - numAdder
	c.client.Set(ctx, "counter", res, 0)
	return nil
}

func (c *CounterRedis) Val() (int, error) {
	val, err := c.client.Get(ctx, "counter").Result()
	if err != nil {
		return 0, err
	}
	if val == "" {
		return 0, errors.New("no value in redis")
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return res, nil
}

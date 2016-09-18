package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// NewPool create a new pooled connection
func NewPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     30,
		MaxActive:   30,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

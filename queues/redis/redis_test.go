package redis

import "testing"

func TestQueue(t *testing.T) {

	redisQueue := NewRedis("192.168.99.100", "6379")
	defer redisQueue.Close()

	redisQueue.Push("test value")
	value, _ := redisQueue.Pop()

	if value != "test value" {
		t.Error("Fetched value is not like expected")
	}
}

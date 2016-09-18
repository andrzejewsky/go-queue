package queues

import (
	"errors"

	"github.com/andrzejewsky/go-queue/queues/redis"
)

//Queue interface for each storages
type Queue interface {
	Push(value string)
	Pop() (string, error)
	Close()
}

// GetQueue get implementation of queue
func GetQueue(name string, config map[string]string) (Queue, error) {
	switch name {
	case "redis":

		host, herr := config["host"]
		port, perr := config["port"]

		if herr && perr {
			return redis.NewRedis(host, port), nil
		}
		return nil, errors.New("Wrong configuration")
	}

	return nil, errors.New("Unsuported implementation")
}

package redis

import "github.com/garyburd/redigo/redis"

//Redis redis implementation
type Redis struct {
	pool *redis.Pool
}

// NewRedis new instance of redis queue
func NewRedis(host string, port string) *Redis {
	pooledConnection := NewPool(host + ":" + port)
	return &Redis{pooledConnection}
}

// Push add value to the queue
func (r *Redis) Push(value string) {
	connection := r.pool.Get()
	defer connection.Close()
	connection.Do("LPUSH", "queue", value)
}

// Pop get and remove value from the queue
func (r *Redis) Pop() (string, error) {
	connection := r.pool.Get()
	defer connection.Close()

	reply, err := redis.Strings(connection.Do("BRPOP", "queue", 0))

	if err != nil {
		return "", err
	}

	return reply[1], nil
}

// Close close connection
func (r *Redis) Close() {
	r.pool.Close()
}

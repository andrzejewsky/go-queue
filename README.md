A simple implementation of queue including basic operations such as Push & Pop.
For now, available implementation has been created only for Redis

Example usage:
```
queue, _ := queues.GetQueue("redis", map[string]string{
  "host": "192.168.99.100",
  "port": "6379",
})

queue.Push("test")
value, _ := queue.Pop()
fmt.Println(value)
defer queue.Close()
```

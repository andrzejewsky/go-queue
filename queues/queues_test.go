package queues

import "testing"

func TestGetQueueWithValidArguments(t *testing.T) {

	_, err := GetQueue("redis", map[string]string{
		"host": "localhost",
		"port": "6379",
	})

	if err != nil {
		t.Error("Instance created corectly but raised error")
	}
}

func TestGetQueueWithMissingConfigParameter(t *testing.T) {

	_, err := GetQueue("redis", map[string]string{})

	if err == nil {
		t.Error("Instance created with error")
	}
}

func TestGetQueueNonExistImplemetation(t *testing.T) {

	_, err := GetQueue("non-exist", map[string]string{})

	if err == nil {
		t.Error("Created non existent instance")
	}
}

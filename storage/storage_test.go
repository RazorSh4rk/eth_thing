package storage

import (
	"sync"
	"testing"
)

func TestStorageSetAndGet(t *testing.T) {
	data = make(map[string]string)

	key := "testKey"
	value := "testValue"

	Set(key, value)

	result := Get(key)
	if result != value {
		t.Errorf("Expected %s, but got %s", value, result)
	}
}

func TestStorageDelete(t *testing.T) {
	data = make(map[string]string)

	key := "testKey"
	value := "testValue"

	Set(key, value)

	Delete(key)

	result := Get(key)
	if result != "" {
		t.Errorf("Expected an empty value, but got %s", result)
	}
}

func TestConcurrentAccess(t *testing.T) {
	data = make(map[string]string)

	// Number of concurrent access
	numAccess := 100
	key := "testKey"
	value := "testValue"

	var wg sync.WaitGroup
	wg.Add(numAccess)

	for i := 0; i < numAccess; i++ {
		go func() {
			defer wg.Done()
			Set(key, value)
			result := Get(key)
			if result != value {
				t.Errorf("Expected %s, but got %s", value, result)
			}
		}()
	}

	wg.Wait()
}

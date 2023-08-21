// Something simimlar to a traditional KV store
// such as redis, so the backend can be swapped
// to something managed.

package storage

import "sync"

// Create a global hashmap and a mutex for thread-safe operations
var data = make(map[string]string)
var dataMutex sync.Mutex

func Set(key string, value string) {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	data[key] = value
}

func Get(key string) string {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	return data[key]
}

func Delete(key string) {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	delete(data, key)
}

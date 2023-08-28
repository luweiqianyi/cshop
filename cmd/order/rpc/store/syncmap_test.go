package store

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	m := sync.Map{}

	// 存储键值对
	m.Store("key1", "value1")
	m.Store("key2", "value2")

	// 覆盖已存在的键值对
	m.Store("key1", "new value1")

	// 输出map中的所有键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}

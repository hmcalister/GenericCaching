package cache

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
	"sync"
)

// Defines a function that returns a cacheable value.
//
// A cachable function always has a single parameter and a single return.
// If your function requires multiple parameters or returns, wrap these in a struct.
type CacheableFunction[ParamType any, ReturnType any] func(ParamType) ReturnType

type Cache[ParamType any, ReturnType any] struct {
	f CacheableFunction[ParamType, ReturnType]

	cacheValues map[uint64]ReturnType

	// Ensure concurrency safety
	concurrentMutex sync.RWMutex
}

func NewCache[ParamType any, ReturnType any](f CacheableFunction[ParamType, ReturnType]) *Cache[ParamType, ReturnType] {
	return &Cache[ParamType, ReturnType]{
		f:           f,
		cacheValues: make(map[uint64]ReturnType),
	}
}


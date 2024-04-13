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

func (c *Cache[ParamType, ReturnType]) CallWithCache(params ParamType) ReturnType {
	paramsHash := c.hashParams(params)

	// Check if the function has been called with these args
	c.concurrentMutex.RLock()
	cachedValue, ok := c.cacheValues[paramsHash]
	c.concurrentMutex.RUnlock()

	if ok {
		return cachedValue
	}

	// We did not hit the cache, so we must compute
	calculatedValue := c.f(params)
	c.concurrentMutex.Lock()
	c.cacheValues[paramsHash] = calculatedValue
	c.concurrentMutex.Unlock()
	return calculatedValue
}

func (c *Cache[ParamType, ReturnType]) hashParams(params ParamType) uint64 {
	var encodedParams bytes.Buffer
	hashFunc := fnv.New64a()
	enc := gob.NewEncoder(&encodedParams)
	if err := enc.Encode(params); err != nil {
		panic(err)
	}
	hashFunc.Write(encodedParams.Bytes())

	// hashFunc := fnv.New64a()
	// pPtr := unsafe.Pointer(&params)
	// pSize := unsafe.Sizeof(params)
	// encodedParams := unsafe.Slice((*byte)(pPtr), pSize)
	// hashFunc.Write(encodedParams)

	return hashFunc.Sum64()
}

package cache

import (
	"testing"
	"time"

	cache "github.com/hmcalister/GenericCaching"
)

func TestCacheStructOfString(t *testing.T) {
	type params struct {
		A string
		B string
	}

	numCalls := 0
	var f = func(p params) string {
		numCalls += 1
		return p.A + p.B
	}

	c := cache.NewCache[params, string](f)

	// Call the function many times with repeated parameters
	// Should only result in one call!
	for i := 0; i < 10; i += 1 {
		c.CallWithCache(params{"A", "B"})
	}

	if numCalls != 1 {
		t.Fatalf("function should have been called once, but was called %v times", numCalls)
	}
}

func TestCacheReturnStruct(t *testing.T) {
	type returnType struct {
		A int
		B int
	}

	numCalls := 0
	var f = func(i int) returnType {
		numCalls += 1
		return returnType{
			A: i,
			B: 10 * i,
		}
	}

	c := cache.NewCache[int, returnType](f)

	// Call the function many times with repeated parameters
	// Should only result in one call!
	for i := 0; i < 10; i += 1 {
		c.CallWithCache(1)
	}

	if numCalls != 1 {
		t.Fatalf("function should have been called once, but was called %v times", numCalls)
	}
}


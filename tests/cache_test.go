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

func TestCacheDirectString(t *testing.T) {
	numCalls := 0
	var f = func(s string) string {
		numCalls += 1
		return s + " cached!"
	}

	c := cache.NewCache[string, string](f)

	// Call the function many times with repeated parameters
	// Should only result in one call!
	for i := 0; i < 10; i += 1 {
		c.CallWithCache("operation?")
	}

	if numCalls != 1 {
		t.Fatalf("function should have been called once, but was called %v times", numCalls)
	}
}

func TestCacheStructOfInteger(t *testing.T) {
	type params struct {
		A int
		B int
	}

	numCalls := 0
	var f = func(p params) int {
		numCalls += 1
		return p.A * p.B
	}

	c := cache.NewCache[params, int](f)

	// Call the function many times with repeated parameters
	// Should only result in one call!
	for i := 0; i < 10; i += 1 {
		c.CallWithCache(params{11, 17})
	}

	if numCalls != 1 {
		t.Fatalf("function should have been called once, but was called %v times", numCalls)
	}
}

func TestCacheDirectInteger(t *testing.T) {
	numCalls := 0
	var f = func(i int) int {
		numCalls += 1
		return i * 2
	}

	c := cache.NewCache[int, int](f)

	// Call the function many times with repeated parameters
	// Should only result in one call!
	for i := 0; i < 10; i += 1 {
		c.CallWithCache(1)
	}

	if numCalls != 1 {
		t.Fatalf("function should have been called once, but was called %v times", numCalls)
	}
}

func TestCacheStructOfPointer(t *testing.T) {
	type indirectParam struct {
		A string
	}

	type params struct {
		I *indirectParam
	}

	numCalls := 0
	var f = func(p params) string {
		numCalls += 1
		return p.I.A
	}

	c := cache.NewCache[params, string](f)
	p := params{I: &indirectParam{"TEST"}}

	// Call the function many times with repeated parameters
	for i := 0; i < 10; i += 1 {
		c.CallWithCache(p)
	}

	if numCalls != 1 {
		t.Fatalf("function should have been called once, but was called %v times", numCalls)
	}

	// Now we change the string on the param and try again
	// If the cache only cared about the pointer, this will not cause additional calls
	// and hence the below test will fail
	p.I.A += " MUTATE"
	for i := 0; i < 10; i += 1 {
		c.CallWithCache(p)
	}

	if numCalls != 2 {
		t.Fatalf("function should have been called twice, but was called %v times", numCalls)
	}
}


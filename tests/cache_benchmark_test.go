package cache_test

import (
	"testing"
	"time"

	cache "github.com/hmcalister/GenericCaching"
)

// ----------------------------------------------------------------------------
// Integer Benchmarks

func BenchmarkCacheDirectIntegerSameParam(b *testing.B) {
	var f = func(i int) int {
		return i * 2
	}
	c := cache.NewCache[int, int](f)

	b.Run("Cache", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			c.CallWithCache(100)
		}
	})

	b.Run("Vanilla", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			f(100)
		}
	})
}

func BenchmarkCacheDirectIntegerVaried(b *testing.B) {
	var f = func(i int) int {
		return i * 2
	}
	c := cache.NewCache[int, int](f)

	b.Run("Cache", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			c.CallWithCache(i % 10)
		}
	})

	b.Run("Vanilla", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			f(i % 10)
		}
	})

}

// ----------------------------------------------------------------------------
// String Benchmarks

func BenchmarkCacheDirectString(b *testing.B) {
	var f = func(s string) int {
		summedChars := 0
		for _, c := range []byte(s) {
			summedChars += int(c)
		}
		return summedChars
	}
	c := cache.NewCache[string, int](f)

	inputString := "The quick brown fox and so on"

	b.Run("Cache", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			c.CallWithCache(inputString)
		}
	})

	b.Run("Vanilla", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			f(inputString)
		}
	})
}

// ----------------------------------------------------------------------------
// Explicitly Expensive Benchmarks

func BenchmarkCacheExpensive(b *testing.B) {
	var f = func(i int) int {
		time.Sleep(time.Microsecond)
		return i * 2
	}
	c := cache.NewCache[int, int](f)

	b.Run("Cache", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			c.CallWithCache(100)
		}
	})

	b.Run("Vanilla", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			f(100)
		}
	})
}

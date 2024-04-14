# Generic Caching

> The two hardest things in computer science: naming things, cache invalidation, and off by one errors

This package implements a caching for functions with a single input and single output. By using structs for function input and output we can cache any general function using struct hashing.

Benchmarks show it is only worth caching for function calls that are relatively expensive. Below function calls taking about 100 nanoseconds simply recomputing the function is less expensive than caching.

Beware! Caching a struct with a `map` field may not work as expected, as the `map` type does not have a guaranteed field order and is hence difficult to hash consistently. In general, do not cache on the `map` type and instead make a call into the `map` from the cached function, using the `map` key as the cache parameter.

Inspired and expanded from [Skarlso/cache](https://github.com/Skarlso/cache).

## Example

See the tests under the `tests` directory for comprehensive examples.

```Go
package main

import (
    "github.com/hmcalister/GenericCaching"
) 

func main() {
    type params struct {
		IntegerParam int
		StringParam  string
		FloatParam   float64
	}

	type output struct {
		StringOutput string
		FloatOutput  float64
	}

	var f = func(p params) output {
		return output{
			StringOutput: p.StringParam,
			FloatOutput:  float64(p.IntegerParam) * p.FloatParam,
		}
	}

	c := cache.NewCache[params, output](f)

    // The result of this call is cached and will be returned 
    // on future calls with the same parameters without recomputation of f
	c.CallWithCache(params{
		IntegerParam: 1,
		StringParam:  "Caching!",
		FloatParam:   2.0,
	})
}
```

## Other Discussions

- Currently, the caching relies on a consistent hashing of the function inputs. This is difficult for structs in general, as discussed above for `map` types. It may be nice to add some consistent caching method for any struct, but this problem pushes up against the language specification.

- The current caching algorithm requires encoding and hashing the parameter type. For generality, we encode the parameters into a string using `encoding/gob` and hash the resulting string using `hash/fnv`. It would be interesting to benchmark just how long this process takes in comparison to just recomputing the function. I imagine for some functions, such a `f(int) -> int` caching would be a detriment.

- In future, it would be nice to add cache types for specifically numerical types only, such that the encoding could be numerical rather than a string, which may be faster.

- In future, adding the option for cache durations would be interesting as well, such that cached values are also associated with a timestamp that allow for recomputation after some time.
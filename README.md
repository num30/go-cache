# go-cache

[![test-and-lint](https://github.com/num30/go-cache/actions/workflows/test.yaml/badge.svg)](https://github.com/num30/go-cache/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/num30/go-cache/branch/main/graph/badge.svg?token=FMvJ4TbC2r)](https://codecov.io/gh/num30/go-cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/num30/go-cache)](https://goreportcard.com/report/github.com/num30/go-cache)
[![Go Reference](https://pkg.go.dev/badge/github.com/num30/go-cache.svg)](https://pkg.go.dev/github.com/num30/go-cache)

ℹ️ go-cache is a port of great [go-cache](https://github.com/patrickmn/go-cache) library that was written by @patrickmn. The main defference is that it relies on generics instead of reflection.

go-cache is a  in-memory key:value store/cache similar to memcached that is
suitable for applications running on a single machine. Its major advantage is
that, being essentially a thread-safe `map[string][T]` with expiration
times, it doesn't need to serialize or transmit its contents over the network.

Any object can be stored, for a given duration or forever, and the cache can be
safely used by multiple goroutines.

Although go-cache isn't meant to be used as a persistent datastore, the entire
cache can be saved to and loaded from a file (using `c.Items()` to retrieve the
items map to serialize, and `NewFrom()` to create a cache from a deserialized
one) to recover from downtime quickly. (See the docs for `NewFrom()` for caveats.)

### Installation

`go get github.com/num30/go-cache`

### Usage

```go
import (
	"fmt"
	"github.com/num30/go-cache"
	"time"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New[string](5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to "lightning", with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", "lightning", cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	// Want performance? Store pointers!
	structCache := cache.New[MyStruct](5*time.Minute, 10*time.Minute)
	structCache.Set("foo", &MyStruct, cache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		fmt.Println(")
	}
}
```

Run this example with `go run examples/main.go`.

## Performance Comparison
Comparison of performance with original [go-cache](https://github.com/patrickmn/go-cache) implementation.

Spoiler alert! The difference is insignificant.



| Test                                         |     [Non generic](https://github.com/patrickmn/go-cache)|       |         This version |           |  
|----------------------------------------------|-----------|-------------|-----------|---------------|
| BenchmarkCacheGetStringExpiring-8            | 29884011 | 41.45 ns/op | 28318041 | 43.31 ns/op |
| BenchmarkCacheGetStringNotExpiring-8         | 91891774 | 14.20 ns/op | 72259294 | 14.02 ns/op |
| BenchmarkCacheGetConcurrentExpiring-8        | 26299849 | 42.58 ns/op | 30129078 | 39.53 ns/op |
| BenchmarkCacheGetConcurrentNotExpiring-8     | 28991383 | 41.25 ns/op | 30760544 | 38.09 ns/op |
| BenchmarkCacheGetManyConcurrentExpiring-8    | 55589712 | 44.90 ns/op | 56991110 | 38.86 ns/op |
| BenchmarkCacheGetManyConcurrentNotExpiring-8 | 30105078 | 43.98 ns/op | 46270045 | 41.54 ns/op |
| BenchmarkCacheSetStringExpiring-8            | 18392893 | 63.41 ns/op | 17788724 | 61.42 ns/op |
| BenchmarkCacheSetStringNotExpiring-8         | 46400654 | 28.45 ns/op | 40226074 | 27.41 ns/op |
| BenchmarkCacheSetDelete-8                    | 18703620 | 60.75 ns/op | 18270448 | 59.90 ns/op |
| BenchmarkCacheSetDeleteSingleLock-8          | 32633755 | 39.34 ns/op | 32415156 | 36.96 ns/op |
| BenchmarkCacheGetStructExpiring-8            | 30487856 | 41.60 ns/op | 26925226 | 40.55 ns/op |
| BenchmarkCacheGetStructNotExpiring-8         | 91921044 | 13.94 ns/op | 96379750 | 13.08 ns/op |
| BenchmarkCacheSetStructExpiring-8            | 13977464 | 86.44 ns/op | 13364509 | 87.69 ns/op |
| BenchmarkCacheSetStructNotExpiring-8         | 22749384 | 54.14 ns/op | 23207397 | 52.58 ns/op |
| BenchmarkCacheSetFatStructExpiring-8         | 11718718 | 103.3 ns/op | 12051895 | 102.3 ns/op |
| BenchmarkCacheGetFatStructNotExpiring-8      | 88695709 | 13.92 ns/op | 83220014 | 13.76 ns/op |


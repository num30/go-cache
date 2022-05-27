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
| <sub>BenchmarkCacheGetStringExpiring-8            </sub>|  <sub>29884011</sub> | <sub>41.45 ns/op</sub> | <sub>28318041</sub> | <sub>43.31 ns/op</sub> |
| <sub>BenchmarkCacheGetStringNotExpiring-8         </sub>|  <sub>91891774</sub> | <sub>14.20 ns/op</sub> | <sub>72259294</sub> | <sub>14.02 ns/op</sub> |
| <sub>BenchmarkCacheGetConcurrentExpiring-8        </sub>|  <sub>26299849</sub> | <sub>42.58 ns/op</sub> | <sub>30129078</sub> | <sub>39.53 ns/op</sub> |
| <sub>BenchmarkCacheGetConcurrentNotExpiring-8     </sub>|  <sub>28991383</sub> | <sub>41.25 ns/op</sub> | <sub>30760544</sub> | <sub>38.09 ns/op</sub> |
| <sub>BenchmarkCacheGetManyConcurrentExpiring-8    </sub>|  <sub>55589712</sub> | <sub>44.90 ns/op</sub> | <sub>56991110</sub> | <sub>38.86 ns/op</sub> |
| <sub>BenchmarkCacheGetManyConcurrentNotExpiring-8 </sub>|  <sub>30105078</sub> | <sub>43.98 ns/op</sub> | <sub>46270045</sub> | <sub>41.54 ns/op</sub> |
| <sub>BenchmarkCacheSetStringExpiring-8            </sub>|  <sub>18392893</sub> | <sub>63.41 ns/op</sub> | <sub>17788724</sub> | <sub>61.42 ns/op</sub> |
| <sub>BenchmarkCacheSetStringNotExpiring-8         </sub>|  <sub>46400654</sub> | <sub>28.45 ns/op</sub> | <sub>40226074</sub> | <sub>27.41 ns/op</sub> |
| <sub>BenchmarkCacheSetDelete-8                    </sub>|  <sub>18703620</sub> | <sub>60.75 ns/op</sub> | <sub>18270448</sub> | <sub>59.90 ns/op</sub> |
| <sub>BenchmarkCacheSetDeleteSingleLock-8          </sub>|  <sub>32633755</sub> | <sub>39.34 ns/op</sub> | <sub>32415156</sub> | <sub>36.96 ns/op</sub> |
| <sub>BenchmarkCacheGetStructExpiring-8            </sub>|  <sub>30487856</sub> | <sub>41.60 ns/op</sub> | <sub>26925226</sub> | <sub>40.55 ns/op</sub> |
| <sub>BenchmarkCacheGetStructNotExpiring-8         </sub>|  <sub>91921044</sub> | <sub>13.94 ns/op</sub> | <sub>96379750</sub> | <sub>13.08 ns/op</sub> |
| <sub>BenchmarkCacheSetStructExpiring-8            </sub>|  <sub>13977464</sub> | <sub>86.44 ns/op</sub> | <sub>13364509</sub> | <sub>87.69 ns/op</sub> |
| <sub>BenchmarkCacheSetStructNotExpiring-8         </sub>|  <sub>22749384</sub> | <sub>54.14 ns/op</sub> | <sub>23207397</sub> | <sub>52.58 ns/op</sub> |
| <sub>BenchmarkCacheSetFatStructExpiring-8         </sub>|  <sub>11718718</sub> | <sub>103.3 ns/op</sub> | <sub>12051895</sub> | <sub>102.3 ns/op</sub> |
| <sub>BenchmarkCacheGetFatStructNotExpiring-8      </sub>|  <sub>88695709</sub> | <sub>13.92 ns/op</sub> | <sub>83220014</sub> | <sub>13.76 ns/op</sub> |


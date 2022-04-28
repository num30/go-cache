# go-cache

ℹ️ go-cache is generic port of [go-cache](https://github.com/patrickmn/go-cache) library.

go-cache is an generic in-memory key:value store/cache similar to memcached that is
suitable for applications running on a single machine. Its major advantage is
that, being essentially a thread-safe `map[string]interface{}` with expiration
times, it doesn't need to serialize or transmit its contents over the network.

Any object can be stored, for a given duration or forever, and the cache can be
safely used by multiple goroutines.

Although go-cache isn't meant to be used as a persistent datastore, the entire
cache can be saved to and loaded from a file (using `c.Items()` to retrieve the
items map to serialize, and `NewFrom()` to create a cache from a deserialized
one) to recover from downtime quickly. (See the docs for `NewFrom()` for caveats.)

### Installation

`go get github.com/DenisPalnitsky/go-cache`

### Usage

```go
import (
	"fmt"
	"github.com/DenisPalnitsky/go-cache"
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

### Reference

`godoc` or [http://godoc.org/github.com/DenisPalnitsky/go-cache](http://godoc.org/github.com/DenisPalnitsky/go-cache)

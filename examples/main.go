package main

import (
	"fmt"
	"time"

	"github.com/num30/go-cache"
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
	structCache := cache.New[*time.Time](5*time.Minute, 10*time.Minute)
	structCache.Set("foo", &time.Time{}, cache.DefaultExpiration)
	if x, found := structCache.Get("foo"); found {
		fmt.Printf("Cached time %+v\n", x)
	}
}

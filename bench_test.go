package cache

type BenchStr struct {
	Num    int
	String string
}

//func Benchmark_CacheGet(b *testing.B) {
//	c := New[*BenchStr](DefaultExpiration, 0)
//	c.Set("foo", &BenchStr{1, "bar"}, DefaultExpiration)
//
//	for i := 0; i < b.N; i++ {
//		c.Get("foo")
//	}
//}

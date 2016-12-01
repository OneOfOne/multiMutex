package multiMutex

import "testing"
import "github.com/OneOfOne/xxhash"

const path = "https://stackoverflow.com/questions/tagged/go"

func BenchmarkModDjb2(b *testing.B) {
	var h uint64
	for i := 0; i < b.N; i++ {
		h += uint64(modDjb2(path))
	}
}
func BenchmarkXXHash(b *testing.B) {
	var h uint64
	for i := 0; i < b.N; i++ {
		h += xxhash.ChecksumString64(path)
	}
}

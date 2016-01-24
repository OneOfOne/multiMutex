package multiMutex

import "testing"

const path = "https://stackoverflow.com/questions/tagged/go"

func BenchmarkModDjb2(b *testing.B) {
	var h int
	for i := 0; i < b.N; i++ {
		h += modDjb2(path)
	}
}

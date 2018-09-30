package test

//执行：go test -v -test.bench .

import (
	"sync"
	"testing"
)

var lock sync.Mutex

func test() {
	lock.Lock()
	lock.Unlock()
}
func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}
func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testdefer()
	}
}



package chanLocker

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChanLocker(t *testing.T) {
	lc := NewLocker()
	result := 0
	for i := 0;i < 1000 ;i++ {
		go func() {
			lc.Lock()
			defer lc.Unlock()
			result ++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(result)
}
func BenchmarkChanLocker(b *testing.B) {
	lc := NewLocker()
	for i := 0; i < b.N; i++ {
		result := 0
		for i := 0;i < 1000 ;i++ {
			go func() {
				lc.Lock()
				defer lc.Unlock()
				result ++
			}()
		}
	}
}

func BenchmarkMutexLocker(b *testing.B) {
	mu := new(sync.Mutex)
	for i := 0; i < b.N; i++ {
		result := 0
		for i := 0;i < 1000 ;i++ {
			go func() {
				mu.Lock()
				defer mu.Unlock()
				result ++
			}()
		}
	}
}


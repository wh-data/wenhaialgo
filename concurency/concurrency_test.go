package concurency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	var (
		n = 5000
	)
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}

	wg.Add(n)
	start := time.Now()
	for i := 0; i < n; i++ {
		go GoPrint(&wg, n)
	}

	wg.Wait()
	span := time.Now().Sub(start).Seconds()
	fmt.Printf("\ntime span %v ", span)
}

func GoPrint(wg *sync.WaitGroup, n int) {
	time.Sleep(time.Millisecond * 800)
	fmt.Printf("%d ", n)
	wg.Done()
}

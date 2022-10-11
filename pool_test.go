package gopool

import (
	"gopool/worker"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

const (
	MaxJob = 10000 // 1w个job
)

var (
	counter1 uint64 = 0
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func BenchmarkWithOneWorker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool := NewPool(1, MaxJob)
		for j := 0; j < MaxJob; j++ {
			pool.AddJob(jobFunc2)
		}
		pool.Wait()
	}
}

func BenchmarkTwoWorker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool := NewPool(2, MaxJob)
		for j := 0; j < MaxJob; j++ {
			pool.AddJob(jobFunc2)
		}
		pool.Wait()
	}
}

func BenchmarkThreeWorker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool := NewPool(3, MaxJob)
		for j := 0; j < MaxJob; j++ {
			pool.AddJob(jobFunc2)
		}
		pool.Wait()
	}
}

func BenchmarkFourWorker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool := NewPool(4, MaxJob)
		for j := 0; j < MaxJob; j++ {
			pool.AddJob(jobFunc2)
		}
		pool.Wait()
	}
}

func BenchmarkTenWorker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool := NewPool(10, MaxJob)
		for j := 0; j < MaxJob; j++ {
			pool.AddJob(jobFunc2)
		}
		pool.Wait()
	}
}

func BenchmarkOneHundredWorker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool := NewPool(100, MaxJob)
		for j := 0; j < MaxJob; j++ {
			pool.AddJob(jobFunc2)
		}
		pool.Wait()
	}
}

func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < MaxJob; j++ {
			jobFunc2()
		}
	}
}

func jobFunc() {
	atomic.AddUint64(&counter1, 1)
}

func jobFunc2() *worker.Result {
	time.Sleep(time.Microsecond)
	return &worker.Result{}
}

/*
BenchmarkWithOneWorker-12                     24          49541128 ns/op
BenchmarkTwoWorker-12                         32          32032489 ns/op
BenchmarkThreeWorker-12                       54          21861044 ns/op
BenchmarkFourWorker-12                        72          16481204 ns/op
BenchmarkTenWorker-12                        150           7967065 ns/op
BenchmarkOneHundredWorker-12                 222           5604295 ns/op
BenchmarkNoPool-12                            25          44055340 ns/op
*/

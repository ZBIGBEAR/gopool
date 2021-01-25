package gopool

import (
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

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func BenchmarkWithOneWorker(b *testing.B) {
	for i:=0;i<b.N;i++{
		pool := NewPool(1, MaxJob)
		pool.Start()
		for j:=0;j<MaxJob;j++{
			pool.AddJob(jobFunc2)
		}
		pool.Stop()
	}
}

func BenchmarkTwoWorker(b *testing.B) {
	for i:=0;i<b.N;i++{
		pool := NewPool(2, MaxJob)
		pool.Start()
		for j:=0;j<MaxJob;j++{
			pool.AddJob(jobFunc2)
		}
		pool.Stop()
	}
}

func BenchmarkThreeWorker(b *testing.B) {
	for i:=0;i<b.N;i++{
		pool := NewPool(3, MaxJob)
		pool.Start()
		for j:=0;j<MaxJob;j++{
			pool.AddJob(jobFunc2)
		}
		pool.Stop()
	}
}

func BenchmarkFourWorker(b *testing.B) {
	for i:=0;i<b.N;i++{
		pool := NewPool(4, MaxJob)
		pool.Start()
		for j:=0;j<MaxJob;j++{
			pool.AddJob(jobFunc2)
		}
		pool.Stop()
	}
}

func BenchmarkTenWorker(b *testing.B) {
	for i:=0;i<b.N;i++{
		pool := NewPool(10, MaxJob)
		pool.Start()
		for j:=0;j<MaxJob;j++{
			pool.AddJob(jobFunc2)
		}
		pool.Stop()
	}
}

func BenchmarkOneHundredWorker(b *testing.B) {
	for i:=0;i<b.N;i++{
		pool := NewPool(100, MaxJob)
		pool.Start()
		for j:=0;j<MaxJob;j++{
			pool.AddJob(jobFunc2)
		}
		pool.Stop()
	}
}

func BenchmarkNoPool(b *testing.B) {
	for i:=0;i<b.N;i++{
		for j:=0;j<MaxJob;j++{
			jobFunc2()
		}
	}
}

func jobFunc(){
	atomic.AddUint64(&counter1,1)
}

func jobFunc2(){
	time.Sleep(time.Microsecond)
}
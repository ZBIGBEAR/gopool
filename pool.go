// master-worker模式
package gopool

import (
	"gopool/master"
	"gopool/worker"
)

const (
	DefaultMaxWorker = 4
	DefaultMaxJob = 1000
)

type Pool struct {
	master *master.Master
	maxWorker int
	maxJob int
}

func newPoolWithMaxWorker(maxWorker int, maxJob int) *Pool {
	return &Pool{
		maxWorker: maxWorker,
		maxJob:    maxJob,
	}
}

func NewDefaultPool() *Pool {
	return newPoolWithMaxWorker(DefaultMaxWorker, DefaultMaxJob)
}

func NewPool(maxWorker int, maxJob int) *Pool {
	return newPoolWithMaxWorker(maxWorker, maxJob)
}

func (p *Pool) Start() {
	p.master = master.NewMaster(p.maxWorker, p.maxJob)
	p.master.Start()
}

func (p *Pool) AddJob(job worker.Job) {
	p.master.AddJob(job)
}

func (p *Pool) Stop() {
	p.master.Stop()
}
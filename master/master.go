package master

import (
	"gopool/worker"
	"sync"
)

type Master struct {
	WorkerQueue     chan *worker.Worker
	jobQueue        chan worker.Job
	quit            chan bool
	quitCollectData chan bool
	ResultQueue     chan *worker.Result
	result          []*worker.Result
	wg              sync.WaitGroup
}

func NewMaster(maxQueue int, maxJob int) *Master {
	return &Master{
		WorkerQueue:     make(chan *worker.Worker, maxQueue),
		jobQueue:        make(chan worker.Job, maxJob),
		quit:            make(chan bool),
		quitCollectData: make(chan bool),
		ResultQueue:     make(chan *worker.Result, maxJob),
		result:          make([]*worker.Result, 0),
	}
}

func (m *Master) Start() {
	// 初始化并启动worker
	for i := 0; i < cap(m.WorkerQueue); i++ {
		worker := worker.NewWorker(m.WorkerQueue, m.ResultQueue)
		worker.Start()
	}

	// 分发任务
	go m.dispatch()

	// 收集结果
	go m.collectResult()
}

func (m *Master) dispatch() {
	for {
		select {
		case job := <-m.jobQueue:
			w := <-m.WorkerQueue
			w.JobChan <- job
		case <-m.quit:
			for i := 0; i < cap(m.WorkerQueue); i++ {
				w := <-m.WorkerQueue
				w.Stop()
			}
		}
	}
}

func (m *Master) collectResult() {
	exit := false
	for {
		if exit {
			break
		}

		select {
		case r := <-m.ResultQueue:
			m.result = append(m.result, r)
			m.wg.Done()
		case <-m.quitCollectData:
			exit = true
		}
	}
}

func (m *Master) stop() {
	m.quit <- true
	m.quitCollectData <- true
}

func (m *Master) AddJob(job worker.Job) {
	m.wg.Add(1)
	m.jobQueue <- job
}

// Wait. 等待队列中的任务全部结束
func (m *Master) Wait() {
	m.wg.Wait()
	m.stop()
}

func (m *Master) Result() []*worker.Result {
	return m.result
}

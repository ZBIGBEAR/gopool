package master

import "gopool/worker"

type Master struct {
	WorkerQueue chan *worker.Worker
	jobQueue chan worker.Job
	quit chan bool
}

func NewMaster(maxQueue int, maxJob int) *Master {
	return &Master{
		WorkerQueue: make(chan *worker.Worker, maxQueue),
		jobQueue:    make(chan worker.Job, maxJob),
		quit:        make(chan bool),
	}
}

func (m *Master) Start(){
	for i :=0;i<cap(m.WorkerQueue);i++{
		worker := worker.NewWorker(m.WorkerQueue)
		worker.Start()
	}
	go m.dispatch()
}

func (m *Master) dispatch(){
	for {
		select {
		case job := <- m.jobQueue:
			w := <- m.WorkerQueue
			w.JobChan <- job
		case <- m.quit:
			for i := 0; i<cap(m.WorkerQueue);i++ {
				w :=  <- m.WorkerQueue
				w.Start()
			}
		}
	}
}

func (m *Master) Stop(){
	m.quit <- true
}

func (m *Master) AddJob(job worker.Job) {
	m.jobQueue <- job
}
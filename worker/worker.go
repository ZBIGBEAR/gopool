package worker

type Job func() *Result

type Worker struct {
	WorkerQueue chan *Worker
	JobChan     chan Job
	quit        chan bool
	ResultQueue chan *Result
}

func NewWorker(workerQueue chan *Worker, resultQueue chan *Result) *Worker {
	return &Worker{
		WorkerQueue: workerQueue,
		ResultQueue: resultQueue,
		JobChan:     make(chan Job),
		quit:        make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w
			select {
			case job := <-w.JobChan:
				w.ResultQueue <- job()
			case <-w.quit:
				break
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.quit <- true
}

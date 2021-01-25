package worker

type Job func()

type Worker struct {
	WorkerQueue chan *Worker
	JobChan chan Job
	quit chan bool
}

func NewWorker(workerQueue chan *Worker) *Worker {
	return &Worker{
		WorkerQueue: workerQueue,
		JobChan:     make(chan Job),
		quit:make(chan bool),
	}
}

func (w *Worker) Start(){
	go func() {
		for {
			w.WorkerQueue <- w
			select {
			case job := <-w.JobChan:
				job()
			case <-w.quit:
				break
			}
		}
	}()
}

func (w *Worker) Stop(){
	w.quit <- true
}
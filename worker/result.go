package worker

type Result struct {
	data interface{}
	err  error
}

func NewResult(data interface{}, err error) *Result {
	return &Result{
		data: data,
		err:  err,
	}
}

func (r *Result) Data() interface{} {
	return r.data
}

func (r *Result) Err() error {
	return r.err
}

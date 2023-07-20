package rpc

// Worker responsible for queue serving.
type Worker struct {
	Queue *Queue
}

// NewWorker initializes a new Worker.
func NewWorker(queue *Queue) *Worker {
	return &Worker{
		Queue: queue,
	}
}

// DoWork processes jobs from the queue (jobs channel).
func (w *Worker) DoWork() bool {
	for {
		select {
		// if context was canceled.
		case <-w.Queue.ctx.Done():
			Logger.Info().Msgf("Work done in queue %s: %s!", w.Queue.name, w.Queue.ctx.Err())
			return true
		// if job received.
		case job := <-w.Queue.jobs:
			err := job.Run()
			if err != nil {
				Logger.Error().Msgf("Error: %v", err)
				continue
			}
		}
	}
}

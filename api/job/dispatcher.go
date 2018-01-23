package job

import (
	"log"
	"time"

	"github.com/quintilesims/layer0/api/lock"
	"github.com/quintilesims/layer0/common/models"
)

const (
	// dynamodb allows a burst-read/write every 5 minutes
	// matching that value here to try and avoid hitting that limit
	dispatcherPeriod = time.Minute * 5
)

func RunWorkersAndDispatcher(numWorkers int, store Store, runner Runner, lock lock.Lock) (*time.Ticker, func()) {
	quitFNs := make([]func(), numWorkers)
	queue := make(chan string)
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i+1, store, queue, runner, lock)
		quitFNs[i] = worker.Start()
	}

	dispatcher := NewDispatcher(store, queue)
	ticker := time.NewTicker(dispatcherPeriod)
	go func() {
		for range ticker.C {
			log.Printf("[INFO] [JobDispatcher] Starting dispatcher")
			if err := dispatcher.Run(); err != nil {
				log.Printf("[ERROR] [JobDispatcher] Failed to dispatch: %v", err)
			}
		}
	}()

	// todo: this feels like a hacky solution; should investigate alternative
	store.SetInsertHook(func(jobID string) {
		go func() { queue <- jobID }()
	})

	quitFN := func() {
		for _, quitFN := range quitFNs {
			quitFN()
		}
	}

	return ticker, quitFN
}

type Dispatcher struct {
	store Store
	queue chan<- string
}

func NewDispatcher(store Store, queue chan<- string) *Dispatcher {
	return &Dispatcher{
		store: store,
		queue: queue,
	}
}

func (d *Dispatcher) Run() error {
	jobs, err := d.store.SelectAll()
	if err != nil {
		return err
	}

	for _, job := range jobs {
		if models.JobStatus(job.Status) == models.PendingJobStatus {
			d.queue <- job.JobID
		}
	}

	return nil
}
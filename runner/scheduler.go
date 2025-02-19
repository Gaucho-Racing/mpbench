package runner

import (
	"mpbench/model"
	"mpbench/utils"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var Queue *Scheduler

func InitializeScheduler(workers string) {
	numWorkers, err := strconv.Atoi(workers)
	if err != nil {
		numWorkers = 10
		utils.SugarLogger.Infof("Invalid number of workers, defaulting to %d", numWorkers)
	}
	Queue = NewScheduler(numWorkers)
	go Queue.Run()
}

type Scheduler struct {
	queue      chan model.Run
	workers    int
	stopChan   chan struct{}
	inProgress int32
}

func NewScheduler(workers int) *Scheduler {
	s := &Scheduler{
		queue:      make(chan model.Run),
		workers:    workers,
		stopChan:   make(chan struct{}),
		inProgress: 0,
	}
	go s.printStats()
	return s
}

func (s *Scheduler) printStats() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			utils.SugarLogger.Infof("Worker Stats: %d runs in progress, %d runs in queue",
				atomic.LoadInt32(&s.inProgress),
				len(s.queue))
		case <-s.stopChan:
			return
		}
	}
}

func (s *Scheduler) Add(run model.Run) {
	s.queue <- run
}

func (s *Scheduler) Run() {
	var wg sync.WaitGroup
	for i := 0; i < s.workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case run, ok := <-s.queue:
					if !ok {
						return
					}
					atomic.AddInt32(&s.inProgress, 1)
					StartRun(run)
					atomic.AddInt32(&s.inProgress, -1)
				case <-s.stopChan:
					return
				}
			}
		}()
	}
	utils.SugarLogger.Infof("Started %d workers", s.workers)
	wg.Wait()
}

func (s *Scheduler) Stop() {
	close(s.stopChan)
	close(s.queue)
}

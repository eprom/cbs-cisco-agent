package task

import (
	"fmt"
	"github.com/jakhog/cbs-cisco-agent/log"
	"sync"
	"time"
)

const DEFAULT_MAX_TICK = time.Minute * 10

type TaskRunner struct {
	task        Task
	initialTick time.Duration
	maxTick     time.Duration
	tick        time.Duration
	failCount   int
	log         *log.Logger
	timer       *time.Timer
	running     bool
	stop        chan struct{}
	immediate   chan struct{}
}

func NewRunnerWithMax(task Task, tick, maxTick time.Duration) *TaskRunner {
	return &TaskRunner{
		task:        task,
		initialTick: tick,
		tick:        tick,
		maxTick:     maxTick,
		running:     false,
		log:         log.NewLogger("Task " + task.Name()),
		immediate:   make(chan struct{}, 1),
	}
}

func NewRunner(task Task, tick time.Duration) *TaskRunner {
	return NewRunnerWithMax(task, tick, DEFAULT_MAX_TICK)
}

func (runner *TaskRunner) Start(wg *sync.WaitGroup) {
	if runner.running {
		return
	}
	// Spawn a new goroutine to run this task
	runner.running = true
	wg.Add(1)
	runner.failCount = 0
	runner.tick = runner.initialTick
	runner.timer = time.NewTimer(runner.tick) // This will give us the first tick
	runner.stop = make(chan struct{})
	go func() {
		// Cleanup for when we are done
		defer func() {
			runner.log.Info("Stopped")
			runner.running = false
			wg.Done()
		}()
		// Run this task forever
		runner.log.Info("Started")
	loop:
		for {
			select {
			case <-runner.stop:
				// A stop is requested, bail out
				break loop
			case <-runner.immediate:
				if !runner.timer.Stop() {
					// We don't want the normal to trigger just after we are done
					<-runner.timer.C
				}
				runner.doRunSafe(false)
			case <-runner.timer.C:
				// Normal run on ticks
				runner.doRunSafe(true)
			}
		}
	}()
}

func (runner *TaskRunner) RunImmediately() {
	select {
	case runner.immediate <- struct{}{}:
	default:
	}
}

func (runner *TaskRunner) Stop() {
	if !runner.running {
		return
	}
	// If we try to stop multiple times
	defer func() { recover() }()
	close(runner.stop)
}

func (runner *TaskRunner) doRunSafe(shouldBackOff bool) {
	// Do the actual work
	err := runner.runSafe()
	if err != nil {
		runner.failCount++
		runner.log.Errorf("%v (has failed %v times)", err, runner.failCount)
		if shouldBackOff {
			// Backof the timer if an error occurs (on scheduled runs)
			if runner.tick < runner.maxTick {
				runner.tick *= 2
				if runner.tick > runner.maxTick {
					runner.tick = runner.maxTick
				}
				runner.log.Info("Increasing tick to", runner.tick)
			}
		}
	} else if runner.failCount > 0 {
		runner.log.Info("Recovered from failures")
		// Make sure we reset tick if an error has occured before
		runner.tick = runner.initialTick
		runner.failCount = 0
	}
	runner.timer.Reset(runner.tick)
	// We don't want to re-run immediately, so drain channel
	select {
	case <-runner.immediate:
	default:
	}
}

func (runner *TaskRunner) runSafe() (err error) {
	// This method can never panic, that would break everything
	defer func() {
		if pan := recover(); pan != nil {
			switch e := pan.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("Caught panic: %v", e)
			}
		}
	}()
	// Do the actual task, and check how long it takes
	start := time.Now()
	runner.task.Run(runner.log)
	took := time.Now().Sub(start)
	// Warn if it takes more than it's own tick
	if took > runner.initialTick {
		runner.log.Warningf("Took %v (should run every %v)", took, runner.initialTick)
	}
	return nil
}

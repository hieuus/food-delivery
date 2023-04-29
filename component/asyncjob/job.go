package asyncjob

import (
	"context"
	"time"
)

// Interface
type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

// Job Struct
type job struct {
	config     jobConfig
	handle     JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

// State
type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

// Job Config
type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

const (
	defaultMaxTimeout    = time.Second * 10
	defaultMaxRetryCount = 3
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 5, time.Second * 10}
)

// Job Handler
type JobHandler func(ctx context.Context) error

// Implement
func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	var err error
	err = j.handle(ctx)

	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted

	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateRetryFailed

	return err
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) RetryIndex() int {
	return j.retryIndex
}

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}
	j.config.Retries = times
}

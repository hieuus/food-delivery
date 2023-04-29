package asyncjob

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	"log"
	"sync"
)

type group struct {
	isConcurrent bool
	jobs         []job
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...job) *group {
	g := &group{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           new(sync.WaitGroup),
	}

	return g
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	for i, _ := range g.jobs {

		if g.isConcurrent {

			go func(aj job) {
				defer common.AppRecover()
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(g.jobs[i])

			continue
		}

		job := g.jobs[i]
		errChan <- g.runJob(ctx, job)
		g.wg.Done()
	}

	g.wg.Wait()

	var err error

	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			err = v
		}
	}

	return err
}

func (g *group) runJob(ctx context.Context, j job) error {
	if err := j.Execute(ctx); err != nil {
		for {
			log.Println(err)

			if j.State() == StateRetryFailed {
				return err
			}

			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}
	return nil
}

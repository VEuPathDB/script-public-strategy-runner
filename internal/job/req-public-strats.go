package job

import (
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

func (j *job) queuePublicStrats() {
	Log().Trace("Start Job.queuePublicStrats")
	defer Log().Trace("End Job.queuePublicStrats")

	strats := j.api.MustGetPublicStrategyList()

	for i := range strats {
		strat := strats[i]
		j.stack.Add(1)
		j.pool.Submit(func() {
			defer j.stack.Done()
			j.copyStrategy(strat)
		})
	}
}

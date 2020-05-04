package job

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/strategy"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
	"time"
)

func (j *job) loadStrategy(id uint, strat *strategy.ShortStrategy, t time.Time) {
	Log().Tracef("Loading strategy copy %s (%d)", strat.Name, id)

	if _, err := j.userApi.GetStrategy(id); err != nil {
		Log().Errorf("Failed to load strategy %s {originalId: %d, copyId: %d}: %s",
			strat.Name, strat.StrategyId, id, err)
		j.stat.Fail++
	} else {
		j.stat.Pass++
	}

	Log().Debugf("Successfully ran public strategy %s (%d) in %s", strat.Name, strat.StrategyId, time.Since(t))
}

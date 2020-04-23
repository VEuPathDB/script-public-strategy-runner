package job

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

func (j *job) loadStrategy(id uint, strat *common.StrategyListingItem) {
	Log().Debugf("Loading strategy copy %s (%d)", strat.Name, id)

	if _, err := j.userApi.GetStrategy(id); err != nil {
		Log().Errorf("Failed to load strategy %s {originalId: %d, copyId: %d}: %s",
			strat.Name, strat.StrategyId, id, err)
		j.stat.Fail++
	} else {
		j.stat.Pass++
	}
}

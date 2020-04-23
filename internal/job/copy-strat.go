package job

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"

	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

func (j *job) copyStrategy(s common.StrategyListingItem) {
	Log().Debugf("Submitting strategy copy request for strategy %s", s.Name)

	res, err := j.userApi.CopyStrategy(s.Signature)
	if err != nil {
		Log().Errorf("Failed to copy strategy %s (%d): %s", s.Name, s.StrategyId, err)
		j.stat.Fail++
		return
	}

	j.stack.Add(1)
	j.pool.Submit(func() {
		defer j.stack.Done()
		j.loadStrategy(res.Id, &s)
	})
}

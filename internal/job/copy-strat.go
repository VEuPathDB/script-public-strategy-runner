package job

import (
	err2 "github.com/VEuPathDB/lib-go-wdk-api/v0/err"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"

	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

func (j *job) copyStrategy(s common.StrategyListingItem) {
	Log().Debugf("Submitting strategy copy request for strategy %s", s.Name)

	res, err := j.userApi.CopyStrategy(s.Signature)
	if err != nil {
		Log().Errorf("Failed to copy strategy %s (%d): %s", s.Name, s.StrategyId, err)
		if tmp, ok := err.(err2.HttpRequestError); ok {
			if tmp.ResponseCode().Exists() && tmp.ResponseCode().Get() == 400 {
				j.stat.Warn++
				return
			}
		}
		j.stat.Fail++
		return
	}

	j.stack.Add(1)
	j.pool.Submit(func() {
		defer j.stack.Done()
		j.loadStrategy(res.Id, &s)
	})
}

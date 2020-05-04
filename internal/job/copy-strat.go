package job

import (
	"time"

	err2 "github.com/VEuPathDB/lib-go-wdk-api/v0/except"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/strategy"

	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

func (j *job) copyStrategy(s strategy.ShortStrategy) {
	Log().Tracef("Submitting strategy copy request for strategy %s", s.Name)
	start := time.Now()

	res, err := j.userApi.CopyStrategy(s.Signature)
	if err != nil {
		Log().Errorf("Failed to copy strategy %s (%d): %s", s.Name, s.StrategyId, err)
		if tmp, ok := err.(err2.HttpRequestError); ok {
			if tmp.ResponseCode() != nil && *tmp.ResponseCode() == 400 {
				j.stat.Warn++
				return
			}
		}
		j.stat.Fail++
		return
	}

	j.loadStrategy(res.Id, &s, start)
}

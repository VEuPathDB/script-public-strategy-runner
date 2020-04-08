package job

import (
	. "github.com/Foxcapades/Go-ChainRequest/simple"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
	"strconv"
)

func (j *job) loadStrategy(id uint) {
	Log().Debugf("Loading strategy %d", id)
	res := GetRequest(j.url.CurrentUserStrategiesUrl() + strconv.Itoa(int(id))).
		Submit()

	if res.GetError() != nil {
		Log().Errorf("Failed to load strategy %d: %s", id, res.GetError())
		return
	}

	if res.MustGetResponseCode() != 200 {
		Log().Warnf("Unexpected response from strategy service for strategy %d: %d",
			id, res.MustGetResponseCode())
		Log().Debug(string(res.MustGetBody()))
	}
}

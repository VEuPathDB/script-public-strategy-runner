package job

import (
	"encoding/json"
	"fmt"
	"github.com/Foxcapades/Go-ChainRequest/request/header"
	. "github.com/Foxcapades/Go-ChainRequest/simple"
	"github.com/VEuPathDB/lib-go-rest-types/veupath/service"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

type copyRes struct {
	Id uint `json:"id"`
}

func (j *job) copyStrategy(s service.Strategy) {
	Log().Debugf("Submitting strategy copy request for strategy %s", s.Name)
	var result copyRes
	res := PostRequest(j.url.CurrentUserStrategiesUrl()).
		SetHeader(header.CONTENT_TYPE, "application/json").
		SetBody(stratCopyBody(s.Signature)).
		Submit()

	if res.GetError() != nil {
		Log().Errorf("Failed to copy strategy %s: %s", s.Name, res.GetError())
		return
	}

	if res.MustGetResponseCode() != 200 {
		Log().Warnf("Unexpected response from strategy service while attempting to copy strategy %s: %d", s.Name, res.MustGetResponseCode())
		Log().Debug(string(res.MustGetBody()))
		return
	}

	if err := res.UnmarshalBody(&result, UnmarshallerFunc(json.Unmarshal)); err != nil {
		Log().Errorf("Could not parse response body while copying strategy %s: %s",
			s.Name, err)
		return
	}

	j.stack.Add(1)
	j.pool.Submit(func() {
		defer j.stack.Done()
		j.loadStrategy(result.Id)
	})
}

func stratCopyBody(s string) []byte {
	return []byte(fmt.Sprintf(`{"sourceStrategySignature":"` + s + `"}`))
}

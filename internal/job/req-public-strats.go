package job

import (
	"encoding/json"
	. "github.com/Foxcapades/Go-ChainRequest/simple"
	"github.com/VEuPathDB/lib-go-rest-types/veupath/service"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

func (j *job) queuePublicStrats() {
	Log().Trace("Start Job.queuePublicStrats")
	defer Log().Trace("End Job.queuePublicStrats")

	var strats []service.Strategy
	url := j.url.PublicStrategiesUrl()

	Log().Debug("Sending get request to", url)
	GetRequest(url).
		Submit().
		MustUnmarshalBody(&strats, UnmarshallerFunc(json.Unmarshal))

	for i := range strats {
		sig := strats[i]
		j.stack.Add(1)
		j.pool.Submit(func() {
			defer j.stack.Done()
			j.copyStrategy(sig)
		})
	}
}

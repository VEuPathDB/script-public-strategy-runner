package job

import (
	"github.com/VEuPathDB/script-public-strategy-runner/internal/out"
	"sync"

	"github.com/VEuPathDB/lib-go-wdk-api/v0"
	wp "github.com/gammazero/workerpool"

	"github.com/VEuPathDB/script-public-strategy-runner/internal/conf"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

func New(conf conf.Configuration, api wdk.Api) Job {
	return &job{
		conf:    conf,
		pool:    wp.New(int(conf.Threads)),
		api:     api,
		userApi: api.CurrentUserApi(),
	}
}

type Job interface {
	Run() out.Summary
}

type job struct {
	stat  out.Summary
	conf  conf.Configuration
	pool  *wp.WorkerPool
	stack sync.WaitGroup

	api     wdk.Api
	userApi wdk.UserApi
}

func (j *job) Run() out.Summary {
	Log().Trace("Start Job.Run()")
	defer Log().Trace("End Job.Run()")
	j.stat.Url = j.api.GetUrl().String()
	j.queuePublicStrats()
	j.stack.Wait()
	j.pool.StopWait()
	return j.stat
}

package job

import (
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
	Run()
}

type job struct {
	conf  conf.Configuration
	pool  *wp.WorkerPool
	stack sync.WaitGroup

	api     wdk.Api
	userApi wdk.UserApi
}

func (j *job) Run() {
	Log().Trace("Start Job.Run()")
	defer Log().Trace("End Job.Run()")
	j.queuePublicStrats()
	j.stack.Wait()
	j.pool.StopWait()
}

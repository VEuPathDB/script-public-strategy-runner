package job

import (
	"github.com/VEuPathDB/lib-go-rest-types/veupath"
	"github.com/VEuPathDB/script-public-strategy-runner/internal/conf"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
	wp "github.com/gammazero/workerpool"
	"sync"
)

func New(conf conf.Configuration) Job {
	return &job{
		conf: conf,
		pool: wp.New(int(conf.Threads)),
		url:  veupath.NewApiUrlBuilder(conf.SiteUrl),
	}
}

type Job interface {
	Run()
}

type job struct {
	conf  conf.Configuration
	pool  *wp.WorkerPool
	stack sync.WaitGroup

	url veupath.ApiUrlBuilder
}

func (j *job) Run() {
	Log().Trace("Start Job.Run()")
	defer Log().Trace("End Job.Run()")
	j.url.SetAuthTkt(j.conf.Auth)
	j.queuePublicStrats()
	j.stack.Wait()
	j.pool.StopWait()
}

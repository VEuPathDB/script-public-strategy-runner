package main

import (
	. "github.com/Foxcapades/Argonaut/v0"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/conf"
	"github.com/VEuPathDB/script-public-strategy-runner/internal/job"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
	"time"
)

var timeStart time.Time

func main() {
	defer func() { Log().Infof("Executed in %s\n", time.Now().Sub(timeStart)) }()
	Log().Info("Starting public strategy runner")
	var config Configuration

	NewCommand().
		Description(AppDescription).
		Flag(NewFlag().
			Short('v').
			Long("verbose").
			Description(VerboseFlagHelp).
			BindUseCount(&config.Verbose)).
		Flag(NewFlag().
			Description(ThreadFlagHelp).
			Short('t').
			Long("threads").
			Bind(&config.Threads, true)).
		Flag(NewFlag().
			Description(AuthFlagHelp).
			Short('a').
			Long("auth").
			Bind(&config.Auth, true)).
		Arg(NewArg().Description("WDK Site url").
			Require().
			TypeHint("url-string").
			Bind(&config.SiteUrl).
			Name("Site URL")).
		MustParse()

	ValidateConfig(&config)

	job.New(config).Run()
}

func init() {
	timeStart = time.Now()
}

package main

import (
	"fmt"
	"os"
	"time"

	. "github.com/Foxcapades/Argonaut/v0"
	"github.com/Foxcapades/Argonaut/v0/pkg/argo"
	"github.com/VEuPathDB/lib-go-wdk-api/v0"

	. "github.com/VEuPathDB/script-public-strategy-runner/internal/conf"
	"github.com/VEuPathDB/script-public-strategy-runner/internal/job"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

var version string
var timeStart time.Time

func main() {
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
		Flag(NewFlag().
			Description("Prints current version").
			Long("version").
			OnHit(func(argo.Flag) {
				fmt.Println(version)
				os.Exit(0)
			})).
		Arg(NewArg().Description("WDK Site url").
			Require().
			TypeHint("url-string").
			Bind(&config.SiteUrl).
			Name("Site URL")).
		MustParse()

	defer func() { Log().Infof("Executed in %s\n", time.Now().Sub(timeStart)) }()
	Log().Info("Starting public strategy runner")

	ValidateConfig(&config)

	job.New(config, wdk.ForceNew(config.SiteUrl).UseAuthToken(config.Auth)).Run()
}

func init() {
	timeStart = time.Now()
}

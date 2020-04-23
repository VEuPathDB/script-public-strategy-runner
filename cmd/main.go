package main

import (
	"encoding/json"
	"fmt"
	"github.com/VEuPathDB/script-public-strategy-runner/internal/out"
	"gopkg.in/yaml.v3"
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
		Flag(SlFlag('a', "auth", AuthFlagHelp).
			Bind(&config.Auth, true)).
		Flag(LFlag("summary", SummaryFlagHelp).
			Arg(NewArg().Name("json|yaml").Bind(&config.Summary).Require())).
		Flag(SlFlag('t', "threads", ThreadFlagHelp).
			Bind(&config.Threads, true)).
		Flag(SlFlag('v', "verbose", VerboseFlagHelp).
			BindUseCount(&config.Verbose)).
		Flag(SlFlag('V', "version", "Prints current version").
			OnHit(func(argo.Flag) {
				fmt.Println(version)
				os.Exit(0)
			})).
		Arg(NewArg().Description("WDK Site url").
			Require().
			Name("url-string").
			Bind(&config.SiteUrl)).
		MustParse()

	defer func() { Log().Infof("Executed in %s\n", time.Now().Sub(timeStart)) }()
	Log().Info("Starting public strategy runner")

	ValidateConfig(&config)

	stats := job.New(config, wdk.ForceNew(config.SiteUrl).UseAuthToken(config.Auth)).
		Run()

	switch config.Summary {
	case out.SummaryTypeJson:
		if e := json.NewEncoder(os.Stdout).Encode(stats); e != nil {
			panic(e)
		}
	case out.SummaryTypeYaml:
		if e := yaml.NewEncoder(os.Stdout).Encode(stats); e != nil {
			panic(e)
		}
	}
}

func init() {
	timeStart = time.Now()
}

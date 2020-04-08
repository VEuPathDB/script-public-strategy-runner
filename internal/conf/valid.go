package conf

import (
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/site"
)

func ValidateConfig(conf *Configuration) {
	if conf.Verbose == 1 {
		Log().LogLevel(LevelDebug)
	} else if conf.Verbose > 1 {
		Log().LogLevel(LevelTrace)
	}

	Log().Tracef("Begin conf.ValidateConfig <- %s", conf)
	defer Log().Tracef("End conf.ValidateConfig -> %s", conf)

	if tmp, err := ResolveUrl(conf.SiteUrl); err != nil {
		panic(err)
	} else {
		conf.SiteUrl = tmp
		Log().Infof("Site URL resolved to: %s", tmp)
	}

	if conf.Threads == 0 {
		conf.Threads = 20
	}
}

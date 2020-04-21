package conf

import (
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

func ValidateConfig(conf *Configuration) {
	if conf.Verbose == 1 {
		Log().LogLevel(LevelDebug)
	} else if conf.Verbose > 1 {
		Log().LogLevel(LevelTrace)
	}

	if conf.Threads == 0 {
		conf.Threads = 20
	}
}

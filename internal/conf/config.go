package conf

import "github.com/VEuPathDB/script-public-strategy-runner/internal/out"

type Configuration struct {
	Verbose int
	SiteUrl string
	Threads uint8
	Auth    string
	Summary out.SummaryType
}

package out

import "fmt"

const (
	errBadST = `unrecognized summary format "%s".  Must be one of (%s, %s)`
)

type SummaryType string

const (
	SummaryTypeJson SummaryType = "json"
	SummaryTypeYaml SummaryType = "yaml"
)

func (s *SummaryType) Unmarshal(value string) error {
	tmp := SummaryType(value)
	if tmp == SummaryTypeJson || tmp == SummaryTypeYaml {
		*s = tmp
		return nil
	}
	return fmt.Errorf(errBadST, tmp, SummaryTypeJson, SummaryTypeYaml)
}

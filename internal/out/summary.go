package out

import (
	"encoding/json"
)

type Summary struct {
	Url string `json:"url" yaml:"URL"`
	Pass uint `json:"pass" yaml:"Pass"`
	Fail uint `json:"500" yaml:"500"`
	Warn uint `json:"400" yaml:"400"`
}

type sumAlias Summary

func (s Summary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		S sumAlias `json:"summary"`
	}{sumAlias(s)})
}

func (s Summary) MarshalYAML() (interface{}, error) {
	return struct {
		S sumAlias `yaml:"Summary"`
	}{sumAlias(s)}, nil
}

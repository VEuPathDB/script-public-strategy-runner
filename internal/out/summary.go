package out

import (
	"encoding/json"
)

type Summary struct {
	Url string `json:"url" yaml:"URL"`
	Pass uint `json:"pass" yaml:"Pass"`
	Fail uint `json:"fail" yaml:"Fail"`
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

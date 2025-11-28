package jsonutils

import (
	"encoding/json"

	"github.com/wudeyong/goutils/log"
)

func ToJson(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func FromJson[T any](param string) *T {
	var t T
	err := json.Unmarshal([]byte(param), &t)
	if err != nil {
		log.Info("Error parse json: {}", err)
		return nil
	}
	return &t
}

package server

import "fmt"

type InputData map[string]interface{}

func (t InputData) String(key string) string {
	if val, ok := t[key]; ok && val != nil {
		return fmt.Sprintf("%v", val)
	}
	return ""
}

func (t InputData) Int64(key string) int64 {
	return 0
}

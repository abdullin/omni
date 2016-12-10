package spec

import (
	"encoding/json"
	"fmt"
)

func unmarshal(d []byte, i interface{}) {
	err := json.Unmarshal(d, i)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal '%s': %s", string(d), err))
	}
}
func marshal(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err != nil {
		panic("Failed to marshal: " + err.Error())
	}
	return b
}
func marshalIndent(i interface{}) []byte {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic("Failed to marshal: " + err.Error())
	}
	return b
}

func marshalToMap(i interface{}) map[string]interface{} {
	item := map[string]interface{}{}
	unmarshal(marshal(i), &item)
	return item
}

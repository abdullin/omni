package spec

import "strings"

type Where map[interface{}]string

func (w Where) Map() map[string]string {

	clean := map[string]string{}

	for k, v := range w {
		key := strings.Trim(string(marshal(k)), `"`)
		clean[key] = v
	}

	return clean
}

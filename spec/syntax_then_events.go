package spec

func Events(events ...interface{}) []interface{} {
	if len(events) == 0 {
		return []interface{}{}
	}
	return events

}

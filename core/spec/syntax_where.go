package spec

type Where map[interface{}]string

func (w Where) Map() map[interface{}]string {
	return w
}

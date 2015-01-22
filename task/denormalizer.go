package task

type denormalizer struct {
	store *store
}

func newDenormalizer(s *store) *denormalizer {
	return &denormalizer{s}
}

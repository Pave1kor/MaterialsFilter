package filter

func (s *FilterLists) Get(key string) bool {
	_, ok := s.values[Elements(key)]
	return !ok
}

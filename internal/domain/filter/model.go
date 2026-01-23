package filter

type Elements string
type FilterLists struct {
	listFilters map[string]map[Elements]any
}

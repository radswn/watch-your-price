package websites

type ceneoSearch struct {
}

// New returns new instance of SearchModule with provided websites
func New() (*ceneoSearch, error) {
	search := &ceneoSearch{}
	return search, nil
}

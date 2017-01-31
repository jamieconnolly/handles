package cli

// Args represents a list of arguments.
type Args []string

// Contains returns true if the given argument is
// contained in the list of arguments.
func (a Args) Contains(needle string) bool {
	for _, arg := range a {
		if arg == needle {
			return true
		}
	}
	return false
}

// First returns the first argument (same as a.Index(0)).
func (a Args) First() string {
	return a.Index(0)
}

// Index returns the i'th argument. It returns an empty
// string if the requested argument does not exist.
func (a Args) Index(i int) string {
	if i < 0 || i >= len(a) {
		return ""
	}
	return a[i]
}

package stringList

type StringList []string

// Return true if StringList contains value
func (sl StringList) Contains(value string) bool {
	for _, entry := range sl {
		if entry == value {
			return true
		}
	}
	return false
}

// Add value to StringList only if it does not exist
func (sl *StringList) AddUnique(value string) {
	if !sl.Contains(value) {
		*sl = append(*sl, value)
	}
}

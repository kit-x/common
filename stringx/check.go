package stringx

// ListContain source string in list
func ListContain(list []string, s string) bool {
	for i := range list {
		if list[i] == s {
			return true
		}
	}
	return false
}

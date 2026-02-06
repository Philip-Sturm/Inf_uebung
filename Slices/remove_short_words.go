package slices

// RemoveShortWords entfernt alle Wörter aus dem Slice,
// die kürzer sind als die angegebene Mindestlänge minLen.
// Das Slice wird direkt über einen Pointer verändert.
func RemoveShortWords(words []string, minLen int) []string {
	var sl []string
	for i := 0; i < len(words); i++ {
		if len([]byte(words[i])) >= minLen {
			sl = append(sl, words[i])
		}
	}

	return sl
}

package slices

// CutAfterWord kürzt das Slice so, dass es nur die Wörter
// bis einschließlich des ersten Vorkommens von `target` enthält.
// Falls `target` nicht vorkommt, bleibt das Slice unverändert.
// Das Slice wird über einen Pointer übergeben und direkt verändert.
func CutAfterWord(words []string, target string) []string {
	var cut int = len(words)
	for i := 0; i < len(words); i++ {
		if words[i] == target {
			cut = i+1
		}
	}

	return  words[0:cut]

	// TODO: Funktion implementieren
}

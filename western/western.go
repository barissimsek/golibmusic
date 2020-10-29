package western

var chromatic = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

var scalemap = map[string][]int{
	"major": []int{2, 2, 1, 2, 2, 2, 1}, // Major intervals starting from the root note
	"minor": []int{2, 1, 2, 2, 1, 2, 2}, // Natural minor intervals starting from the root note
}

// Chromatic returns chromatic scale
func Chromatic() []string {
	return chromatic
}

// ScaleFormula returns the scale foruma for a given mode
func ScaleFormula(mode string) []int {
	if _, ok := scalemap[mode]; ok {
		return scalemap[mode]
	}

	return nil
}

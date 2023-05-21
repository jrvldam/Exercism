package strand

var toRNAmap = map[rune]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	result := make([]byte, len(dna))
	for i, nuc := range dna {
		if val, ok := toRNAmap[nuc]; ok {
			result[i] = val
		}
	}

	return string(result)
}

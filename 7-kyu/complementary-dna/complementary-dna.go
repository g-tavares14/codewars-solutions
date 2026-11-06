package kata

func DNAStrand(dna string) string {
	var dnaToRna = map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}

	var rna string

	for _, char := range dna {
		rna += dnaToRna[string(char)]
	}

	return rna
}

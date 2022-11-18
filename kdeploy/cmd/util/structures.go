package cmd

func SliceMapping[I any, O any](inputs []I, mapping func(I) O) (output []O) {
	results := make([]O, len(inputs))
	for i, input := range inputs {
		results[i] = mapping(input)
	}
	return results
}

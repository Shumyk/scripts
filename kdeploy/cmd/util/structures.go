package cmd

func SliceMapping[I any, O any](inputs []I, mapping func(I) O) (output []O) {
	output = make([]O, len(inputs))
	for i, input := range inputs {
		output[i] = mapping(input)
	}
	return
}

func MapToSliceMapping[I any, O any](inputs map[string]I, mapping func(string, I) O) (output []O) {
	output = make([]O, len(inputs))
	position := 0
	for key, value := range inputs {
		output[position] = mapping(key, value)
		position++
	}
	return
}

func ReturnKey[K, V any](key K, value V) K {
	return key
}

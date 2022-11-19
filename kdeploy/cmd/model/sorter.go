package model

import "sort"

func Sorted(options []ImageOption) []ImageOption {
	sort.SliceStable(options, sortByCreated(options))
	return options
}

func sortByCreated(options []ImageOption) func(i, j int) bool {
	return func(i, j int) bool {
		return options[i].Created.After(options[j].Created)
	}
}

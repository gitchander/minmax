package minmax

import "sort"

func _() {
	var a sort.Interface
	var _ Interface = a
}

type Interface interface {
	Len() int
	Less(i, j int) bool
}

func IndexOfMin(v Interface) (index int, ok bool) {
	n := v.Len()
	if n < 1 {
		return 0, false
	}
	index = 0
	for i := 1; i < n; i++ {
		if v.Less(i, index) {
			index = i
		}
	}
	return index, true
}

func IndexOfMax(v Interface) (index int, ok bool) {
	n := v.Len()
	if n < 1 {
		return 0, false
	}
	index = 0
	for i := 1; i < n; i++ {
		if v.Less(index, i) {
			index = i
		}
	}
	return index, true
}

func IndexOfMinMax(v Interface) (indexMin, indexMax int, ok bool) {
	n := v.Len()
	if n < 1 {
		return 0, 0, false
	}
	indexMin, indexMax = 0, 0
	for i := 1; i < n; i++ {
		if v.Less(i, indexMin) {
			indexMin = i
		}
		if v.Less(indexMax, i) {
			indexMax = i
		}
	}
	return indexMin, indexMax, true
}

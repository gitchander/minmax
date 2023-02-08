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

func IndexOfMin(v Interface) (imin int) {
	imin = 0
	n := v.Len()
	for i := 1; i < n; i++ {
		if v.Less(i, imin) {
			imin = i
		}
	}
	return imin
}

func IndexOfMax(v Interface) (imax int) {
	imax = 0
	n := v.Len()
	for i := 1; i < n; i++ {
		if v.Less(imax, i) {
			imax = i
		}
	}
	return imax
}

func IndexOfMinMax(v Interface) (imin, imax int) {
	imin, imax = 0, 0
	n := v.Len()
	for i := 1; i < n; i++ {
		if v.Less(i, imin) {
			imin = i
		}
		if v.Less(imax, i) {
			imax = i
		}
	}
	return imin, imax
}

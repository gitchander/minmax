package minmax

import (
	"golang.org/x/exp/constraints"
)

type Ordered = constraints.Ordered

func Min[T Ordered](as ...T) T {
	n := len(as)
	if n == 0 {
		var zero T
		return zero
	}
	k := 0 // index of min
	for i := 1; i < n; i++ {
		if as[k] > as[i] {
			k = i
		}
	}
	return as[k]
}

func Max[T Ordered](as ...T) T {
	n := len(as)
	if n == 0 {
		var zero T
		return zero
	}
	k := 0 // index of max
	for i := 1; i < n; i++ {
		if as[k] < as[i] {
			k = i
		}
	}
	return as[k]
}

func Min2[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max2[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min3[T Ordered](a, b, c T) T {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func Max3[T Ordered](a, b, c T) T {
	if a > b {
		if a > c {
			return a
		}
	} else {
		if b > c {
			return b
		}
	}
	return c
}

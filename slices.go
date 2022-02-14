package minmax

import (
	"errors"
)

var ErrNoParams = errors.New("minmax: there are no parameters")

//------------------------------------------------------------------------------
type IntSlice []int

func (v IntSlice) Len() int           { return len(v) }
func (v IntSlice) Less(i, j int) bool { return v[i] < v[j] }

var _ Interface = IntSlice([]int{})

func MinInt(vs ...int) int {
	index, ok := IndexOfMin(IntSlice(vs))
	if !ok {
		panic(ErrNoParams)
	}
	return vs[index]
}

func MaxInt(vs ...int) int {
	index, ok := IndexOfMax(IntSlice(vs))
	if !ok {
		panic(ErrNoParams)
	}
	return vs[index]
}

//------------------------------------------------------------------------------
type Float64Slice []float64

func (v Float64Slice) Len() int           { return len(v) }
func (v Float64Slice) Less(i, j int) bool { return v[i] < v[j] }

var _ Interface = Float64Slice([]float64{})

func MinFloat64(vs ...float64) float64 {
	index, ok := IndexOfMin(Float64Slice(vs))
	if !ok {
		panic(ErrNoParams)
	}
	return vs[index]
}

func MaxFloat64(vs ...float64) float64 {
	index, ok := IndexOfMax(Float64Slice(vs))
	if !ok {
		panic(ErrNoParams)
	}
	return vs[index]
}

//------------------------------------------------------------------------------
type StringSlice []string

func (v StringSlice) Len() int           { return len(v) }
func (v StringSlice) Less(i, j int) bool { return v[i] < v[j] }

var _ Interface = StringSlice([]string{})

//------------------------------------------------------------------------------

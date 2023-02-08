package minmax

// ------------------------------------------------------------------------------
type IntSlice []int

func (v IntSlice) Len() int           { return len(v) }
func (v IntSlice) Less(i, j int) bool { return v[i] < v[j] }

var _ Interface = IntSlice([]int{})

func MinInts(vs ...int) int {
	if len(vs) == 0 {
		return 0 // default value
	}
	imin := IndexOfMin(IntSlice(vs))
	return vs[imin]
}

func MaxInts(vs ...int) int {
	if len(vs) == 0 {
		return 0
	}
	imax := IndexOfMax(IntSlice(vs))
	return vs[imax]
}

// ------------------------------------------------------------------------------
type Float64Slice []float64

func (v Float64Slice) Len() int           { return len(v) }
func (v Float64Slice) Less(i, j int) bool { return v[i] < v[j] }

var _ Interface = Float64Slice([]float64{})

func MinFloat64s(vs ...float64) float64 {
	if len(vs) == 0 {
		return 0
	}
	imin := IndexOfMin(Float64Slice(vs))
	return vs[imin]
}

func MaxFloat64s(vs ...float64) float64 {
	if len(vs) == 0 {
		return 0
	}
	imax := IndexOfMax(Float64Slice(vs))
	return vs[imax]
}

// ------------------------------------------------------------------------------
type StringSlice []string

func (v StringSlice) Len() int           { return len(v) }
func (v StringSlice) Less(i, j int) bool { return v[i] < v[j] }

var _ Interface = StringSlice([]string{})

//------------------------------------------------------------------------------

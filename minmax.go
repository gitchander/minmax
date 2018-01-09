package minmax

type Interface interface {
	Len() int
	Less(i, j int) bool
}

func IndexOfMin(v Interface) (min int) {
	n := v.Len()
	if n == 0 {
		return -1
	}
	for i := 1; i < n; i++ {
		if v.Less(i, min) {
			min = i
		}
	}
	return
}

func IndexOfMax(v Interface) (max int) {
	n := v.Len()
	if n == 0 {
		return -1
	}
	for i := 1; i < n; i++ {
		if v.Less(max, i) {
			max = i
		}
	}
	return
}

func IndexOfMinMax(v Interface) (min, max int) {
	n := v.Len()
	if n == 0 {
		return -1, -1
	}
	for i := 1; i < n; i++ {
		if v.Less(i, min) {
			min = i
		}
		if v.Less(max, i) {
			max = i
		}
	}
	return
}

type IntSlice []int

func (v IntSlice) Len() int           { return len(v) }
func (v IntSlice) Less(i, j int) bool { return v[i] < v[j] }

type StringSlice []string

func (v StringSlice) Len() int           { return len(v) }
func (v StringSlice) Less(i, j int) bool { return v[i] < v[j] }

var (
	_ Interface = IntSlice(nil)
	_ Interface = StringSlice(nil)
)

func MinInt(vs ...int) int {
	min := IndexOfMin(IntSlice(vs))
	if min == -1 {
		panic("minmax: MinInt has no parameters")
	}
	return vs[min]
}

func MaxInt(vs ...int) int {
	max := IndexOfMax(IntSlice(vs))
	if max == -1 {
		panic("minmax: MaxInt has no parameters")
	}
	return vs[max]
}

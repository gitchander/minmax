package main

import (
	"fmt"

	"github.com/gitchander/minmax"
)

func main() {
	exampleIntSlice()
	exampleInts()
	exampleNoParameters()
	exampleUseInterface()
}

func exampleIntSlice() {
	ds := []int{1, -2, 0, 7, -5, 3, -1}
	fmt.Println("for slice:", ds)

	vs := minmax.IntSlice(ds)

	min, ok := minmax.IndexOfMin(vs)
	if ok {
		fmt.Println("min:", vs[min])
	}

	max, ok := minmax.IndexOfMax(vs)
	if ok {
		fmt.Println("min:", vs[max])
	}
}

func exampleInts() {
	vs := []int{1, -2, 0, 7, -5, 3, -1}
	fmt.Println("for slice:", vs)
	fmt.Println("min:", minmax.MinInt(vs...))
	fmt.Println("max:", minmax.MaxInt(vs...))
}

func exampleNoParameters() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	fmt.Println("min:", minmax.MinInt())
}

type person struct {
	name string
	age  int
}

type byAge []person

func (p byAge) Len() int           { return len(p) }
func (p byAge) Less(i, j int) bool { return p[i].age < p[j].age }

func exampleUseInterface() {
	vs := []person{
		{name: "Mark", age: 18},
		{name: "Stan", age: 25},
		{name: "Jon", age: 12},
	}
	min, ok := minmax.IndexOfMin(byAge(vs))
	if ok {
		fmt.Println("min:", vs[min])
	}

	max, ok := minmax.IndexOfMax(byAge(vs))
	if ok {
		fmt.Println("max:", vs[max])
	}
}

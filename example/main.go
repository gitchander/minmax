package main

import (
	"fmt"

	"github.com/gitchander/minmax"
)

func main() {
	exampleMax()
	exampleIntSlice()
	exampleInts()
	exampleNoParameters()
	exampleUseInterface()
}

func exampleMax() {
	ds := []int{-47, -70, -13}
	fmt.Println(minmax.MaxInts(ds...))
}

func exampleIntSlice() {
	ds := []int{1, -2, 0, 7, -5, 3, -1}
	fmt.Println("for slice:", ds)

	vs := minmax.IntSlice(ds)

	imin := minmax.IndexOfMin(vs)
	fmt.Println("min:", vs[imin])

	imax := minmax.IndexOfMax(vs)
	fmt.Println("min:", vs[imax])
}

func exampleInts() {
	vs := []int{1, -2, 0, 7, -5, 3, -1}
	fmt.Println("for slice:", vs)
	fmt.Println("min:", minmax.MinInts(vs...))
	fmt.Println("max:", minmax.MaxInts(vs...))
}

func exampleNoParameters() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	fmt.Println("min:", minmax.MinInts())
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
	imin := minmax.IndexOfMin(byAge(vs))
	fmt.Println("min:", vs[imin])

	imax := minmax.IndexOfMax(byAge(vs))
	fmt.Println("max:", vs[imax])
}

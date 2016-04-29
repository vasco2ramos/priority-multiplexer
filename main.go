package main

import (
	"fmt"
	"sort"
)

func allocateByCapacity(resources int, gs Groups) []int {

	gsc := make(Groups, len(gs))
	copy(gsc, gs)
	sort.Sort(ByCapacity{gsc})

	n := len(gsc)
	optimal := resources / n

	result := []int{}

	for _, g := range gsc {
		if g.Capacity >= optimal {
			resources -= optimal
			result = append(result, optimal)
		} else {
			resources -= g.Capacity
			result = append(result, g.Capacity)
		}
		n--
		// HACK: THIS SHOULD BE IN THE FOR CYCLE
		if n != 0 {
			optimal = resources / n
		}
	}
	return result
}

func allocateByPriority(resources int, gs Groups) []int {

	gsc := make(Groups, len(gs))
	copy(gsc, gs)
	sort.Sort(ByPriority{gsc})

	tp := gs.TotalPriority()

	result := []int{}

	for {
		if len(gsc) == 0 || resources == 0 {
			break
		}
		x, gsc := gsc[len(gsc)-1], gsc[:len(gsc)-1]

		optimal := (x.Priority / tp) * resources
		resources -= optimal
		x.Capacity -= optimal

		if x.Capacity <= optimal {
			resources -= x.Capacity
			// calculate result later
		} else {
			x.Capacity -= optimal
			gsc = append(gsc, x)
			// calculate result later
		}

	}
	return result
}

func main() {
	var s = []*Group{
		{
			Name:     "United States",
			Capacity: 250,
			Priority: 2,
		},
		{
			Name:     "Bahamas",
			Capacity: 100,
			Priority: 1,
		},
		{
			Name:     "Japan",
			Capacity: 150,
			Priority: 3,
		},
	}
	fmt.Println(allocateByPriority(50, s))

}

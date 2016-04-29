package main

import (
	"fmt"
	"sort"

	. "priority-multiplexer/model"
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

func allocateByPriority(resources int, gs Groups) Groups {

	gsc := make(Groups, len(gs))
	copy(gsc, gs)
	//sort.Sort(ByPriority{gsc})

	tp := gs.TotalPriority()

	result := Groups{}

	for {
		gsc.Print()
		if len(gsc) == 0 || resources == 0 {
			for _, g := range gsc {
				result = append(result, g)
			}
			break
		}
		x, gsc := gsc[len(gsc)-1], gsc[:len(gsc)-1]

		optimal := int(float64(x.Priority) / float64(tp) * float64(resources))

		if x.Capacity <= optimal {
			print("here1\n")
			resources = resources - x.Capacity
			x.ResourcesAlocated = x.ResourcesAlocated + x.Capacity
			result = append(result, x)
		} else {
			print("here2\n")
			fmt.Println(resources)
			x.Capacity = x.Capacity - optimal
			x.ResourcesAlocated = x.ResourcesAlocated + x.Capacity
			resources -= optimal
			gsc = append(gsc, x)
			resources -= optimal
		}
	}

	return result
}

func main() {
	var s = Groups{
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
	s.Print()
	allocateByPriority(50, s).Print()

}

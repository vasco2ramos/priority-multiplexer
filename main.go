package main

import (
	"sort"
	"fmt"

	. "priority-multiplexer/model"
)

func allocateByCapacity(resources int, gs Groups) []int{

	gsc := make(Groups, len(gs))
	copy(gsc, gs)

	sort.Sort(ByCapacity{gsc})

	n := len(gsc)
  optimal := resources/n

	result := []int{}

  for _, g := range gsc {
		if g.Capacity >= optimal {
				resources -= optimal
				result = append(result,optimal)
		} else{
				resources -=  g.Capacity
				result = append(result, g.Capacity)
		}
		n--
		// HACK: THIS SHOULD BE IN THE FOR CYCLE
		if n != 0 {
			optimal = resources/n
		}
  }
	return result
}


func allocateByPriority(resources int, gs Groups) []int{

	gsc := make(Groups, len(gs))
	copy(gsc, gs)

	sort.Sort(ByPriority{gsc})

	//n := len(gsc)
	//tp := gs.TotalPriority()

	result := []int{}

	for {
		// Give resources to highest priority group
		//p := gsc[0].Priority
		//c := gsc[0].Capacity

		// Decrease group capacity


		// If group still has capacity
		/*
		if c > 0{
			// Put it at the end of the list
		} else { // group still is out of capacity
			// Remove it from the list
		}
*/
		// return division of Resources
	}
	return result
}


func main() {
	var s = []*Group{
	    {
	      Name: "United States",
	      Capacity: 250,
	      Priority: 2,
	    },
	    {
	      Name: "Bahamas",
	      Capacity: 100,
	      Priority: 1,
	    },
	    {
	      Name: "Japan",
	      Capacity: 150,
	      Priority: 3,
	    },
	}
	fmt.Println(allocateByCapacity(50, s))

}

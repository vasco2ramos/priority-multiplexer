package main

import (
	"sort"
	"fmt"

	. "priority-multiplexer/model"
)

func alocateByCapacity(resources int, gs Groups) []int{

		sort.Sort(ByCapacity{gs})

		n := len(gs)
    optimal := resources/n

		result := []int{}

    for _, g := range gs {
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


func alocateByPriority(resources int, gs Groups) []int{
	sort.Sort(ByPriority{gs})

	//var n int = len(gs)

	result := []int{}


	for resources > 0 {



		//optimal := resources/n
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
	fmt.Println(alocateByCapacity(50, s))

}

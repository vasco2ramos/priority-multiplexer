package main

import (
	"sort"
	"math"

	. "priority-multiplexer/model"
)

func alocateByCapacity(resources int, groups []*Group) []int{

		sort.Sort(ByCapacity{groups})

		n := len(groups)
    optimal := resources/n

		result := []int{}

    for i, g := range groups {
			if g.Capacity >= optimal {
					resources -= optimal
					Append(result,optimal)
			} else{
					resources -=  g.Capacity;
					Append(result, g.Capacity);
			}
			n--
			optimal = math.Floor(resources/n);
    }
}


/*
func alocateByPriority(resources int, Groups []*Group) []int{
    numGroups := len(capacity)
    fairValue := resources/numGroups // Everyone gets the same ammount
    // Find real capacity of each group
    for i, c := range capacity {

    }
    // Alocate Resources Between Groups
} */

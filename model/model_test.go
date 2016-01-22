package model_test

import (
	"fmt"
	"sort"

	. "../model"
)

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

func Test() {

    sort.Sort(ByCapacity{s})
    fmt.Println("Groups by Capacity:")
    printGroups(s)

    sort.Sort(ByName{s})
    fmt.Println("Organs by Name:")
    printGroups(s)

    sort.Sort(ByPriority{s})
    fmt.Println("Organs by Name:")
    printGroups(s)
}


func printGroups(s []*Group) {
	for _, o := range s {
		fmt.Printf("%-8s (%d) Priority - (%d)\n", o.Name, o.Capacity, o.Priority)
	}
}

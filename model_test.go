package model_test

var s = []*Group{
    {
      name: "United States",
      capacity: 250,
      priority: 2,
    },
    {
      name: "Bahamas",
      capacity; 100,
      priority: 1,
    },
    {
      name: "Japan",
      capacity: 150,
      priority: 3,
    },
}


func main() {

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

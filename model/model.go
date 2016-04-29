package model

import "fmt"

// Group represents a sample of group to receive resources
type Group struct {
	Name              string
	Capacity          int
	Priority          int
	ResourcesAlocated int
}

// Groups is a representation to make easier the creation of an array of Groups
type Groups []*Group

// Len to calculate the length of Groups
func (s Groups) Len() int { return len(s) }

// Swap swap element in groups
func (s Groups) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByName struct{ Groups }

//Less provides the sort.Interface methods
func (s ByName) Less(i, j int) bool {
	return s.Groups[i].Name < s.Groups[j].Name
}

// ByCapacity implements sort.Interface by providing Less and using the Len and.
type ByCapacity struct{ Groups }

//Less provides the sort.Interface methods
func (s ByCapacity) Less(i, j int) bool {
	return s.Groups[i].Capacity < s.Groups[j].Capacity
}

// ByPriority implements sort.Interface by providing Less and using the Len and
type ByPriority struct{ Groups }

//Less provides the sort.Interface methods
func (s ByPriority) Less(i, j int) bool {
	return s.Groups[i].Priority > s.Groups[j].Priority
}

//TotalPriority gets groups Priority
func (s Groups) TotalPriority() int {
	i := 0
	for _, g := range s {
		i += g.Priority
	}
	return i
}

//TotalPriority gets groups Priority
func (s Groups) Print() {
	for _, g := range s {
		fmt.Printf("%s gets %d\n", g.Name, g.ResourcesAlocated)
	}
}

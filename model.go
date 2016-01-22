package model

import (
	"fmt"
	"sort"
)

type group struct {
    name            string
    capacity        int
    priority        int
}

type groups []*group

func (s groups) Len() int      { return len(s) }
func (s groups) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByName struct{ groups }

func (s ByName) Less(i, j int) bool { return s.groups[i].Name < s.groups[j].Name }

// ByCapacity implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByCapacity struct{ groups }

func (s ByCapacity) Less(i, j int) bool { return s.groups[i].capacity < s.groups[j].capacity }

// ByPriority implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByPriority struct{ groups }

func (s ByPriority) Less(i, j int) bool { return s.groups[i].priority < s.groups[j].priority }

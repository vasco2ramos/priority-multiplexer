package model

type Group struct {
    Name            string
    Capacity        int
    Priority        int
}

type Groups []*Group

func (s Groups) Len() int      { return len(s) }
func (s Groups) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByName struct{ Groups }

func (s ByName) Less(i, j int) bool {
  return s.Groups[i].Name < s.Groups[j].Name
}

// ByCapacity implements sort.Interface by providing Less and using the Len and.
type ByCapacity struct{ Groups }

func (s ByCapacity) Less(i, j int) bool {
  return s.Groups[i].Capacity < s.Groups[j].Capacity
}

// ByPriority implements sort.Interface by providing Less and using the Len and
type ByPriority struct{ Groups }

func (s ByPriority) Less(i, j int) bool {
  return s.Groups[i].Priority < s.Groups[j].Priority
}

func (s Groups) TotalPriority() int {
  i := 0
  for _, g := range s {
    i += g.Priority
  }
  return i
}

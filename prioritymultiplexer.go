package prioritymultiplexer

import (
	"sort"

	. "priority-multiplexer/model"
)

func allocateByCapacity(resources int, gs Groups) Groups {

	gsc := make(Groups, len(gs))
	copy(gsc, gs)
	sort.Sort(ByCapacity{gsc})

	n := len(gsc)
	optimal := resources / n

	result := Groups{}

	for _, g := range gsc {
		if g.Capacity >= optimal {
			resources -= optimal
			g.ResourcesAlocated = optimal
			result = append(result, g)
		} else {
			resources -= g.Capacity
			result = append(result, g)
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
	sort.Sort(ByPriority{gsc})

	tp := gs.TotalPriority()

	result := Groups{}

	for {
		if len(gsc) == 0 || resources == 0 {
			for _, g := range gsc {
				result = append(result, g)
			}
			break
		}
		x := gsc[0]
		gsc = gsc[1:len(gsc)]
		optimal := int(float64(x.Priority) / float64(tp) * float64(resources))

		if x.Capacity <= optimal {
			resources = resources - x.Capacity
			x.ResourcesAlocated = x.ResourcesAlocated + x.Capacity
			result = append(result, x)
		} else {
			if optimal == 0 {
				optimal = resources
			}
			x.Capacity = x.Capacity - optimal
			x.ResourcesAlocated = x.ResourcesAlocated + optimal
			resources -= optimal
			gsc = append(gsc, x)
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
	allocateByCapacity(100, s).Print()
}

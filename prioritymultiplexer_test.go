package prioritymultiplexer_test

import (
	"testing"

	. "priority-multiplexer/model"
)

var t1g = Groups{
	{
		Name:     "United States",
		Capacity: 0,
		Priority: 2,
	},
	{
		Name:     "Japan",
		Capacity: 0,
		Priority: 3,
	},
}

var t1c = 500

var t2g = Groups{
	{
		Name:     "United States",
		Capacity: 100,
		Priority: 2,
	},
	{
		Name:     "Japan",
		Capacity: 200,
		Priority: 3,
	},
}

var t2c = 0

var t3g = Groups{
	{
		Name:     "United States",
		Capacity: 600,
		Priority: 2,
	},
	{
		Name:     "Japan",
		Capacity: 1000,
		Priority: 3,
	},
}

var t3c = 200

func TestAllocateByCapacity(t *testing.T) {
}

func TestAllocateByPriority(t *testing.T) {
}

package model_test

import (
	"bytes"
	"fmt"
	"sort"
	"testing"

	. "priority-multiplexer/model"
)

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

func TestByCapacity(t *testing.T) {
	sort.Sort(ByCapacity{s})
	var result = printGroups(s)
	if result != "Bahamas-100-1Japan-150-3United States-250-2" {
		t.Error("Bad Capacity Sorting")
	}
}

func TestByName(t *testing.T) {
	sort.Sort(ByName{s})
	var result = printGroups(s)
	if result != "Bahamas-100-1Japan-150-3United States-250-2" {
		t.Error("Bad Name Sorting")
	}
}

func TestByPriority(t *testing.T) {
	sort.Sort(ByPriority{s})
	var result = printGroups(s)
	if result != "Bahamas-100-1United States-250-2Japan-150-3" {
		t.Error("Bad Priority Sorting")
	}
}

func TestTotalPriority(t *testing.T) {
	var result = s.TotalPriority()
	if result != 6 {
		t.Error("Wrong Priority")
	}
}

func printGroups(s []*Group) string {
	buffer := bytes.NewBufferString("")

	for _, o := range s {
		fmt.Fprint(buffer, o.Name, "-", o.Capacity, "-", o.Priority)
	}

	return buffer.String()
}

package main

type test struct {
    ammount int
    numOffers int
    prio []int
    expected int
}

var tests = []test{
    {
        ammount: 100,
        numOffers: 1,
        prio: []int{1},
        expected: 100,
    },
    {
        ammount: 200,
        numOffers: 1,
        prio: []int{1},
        expected: 200,
    },
    {
        ammount: 300,
        numOffers: 1,
        prio: []int{1},
        expected: 300,
    },

}
func main(){
    testAll()
}

func testAll() {
    for i, t := range tests {
        output := f(t.ammount, nil, nil)
        if output != t.expected{
            fmt.Printf("%d) Boom\n", i+1)
        }
    }
}


func f(resources, capacity []int, priority []int) int{
    numGroups := len(capacity)
    fairValue := resources/numGroups // Everyone gets the same ammount
    // Find real capacity of each group
    for i, c := range capacity {

    }
    // Alocate Resources Between Groups
}

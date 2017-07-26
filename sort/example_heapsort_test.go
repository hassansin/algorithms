package sort_test

import (
	"fmt"

	pq "github.com/hassansin/algorithms/priorityqueue"
	"github.com/hassansin/algorithms/sort"
)

type Person struct {
	Name string
	Age  int
}

func (d1 Person) Less(d pq.Key) bool {
	d2 := d.(Person)
	return d1.Age-d2.Age < 0
}

func ExampleHeapsort() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	data := make([]pq.Key, len(people))
	for i := range people {
		data[i] = pq.Key(people[i])
	}

	fmt.Println(data)
	sort.Heapsort(data)
	fmt.Println(data)

	// Output:
	// [{Bob 31} {John 42} {Michael 17} {Jenny 26}]
	// [{Michael 17} {Jenny 26} {Bob 31} {John 42}]
}

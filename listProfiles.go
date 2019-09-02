package testgopkg

import "fmt"

type Profile struct {
	Name string
	Age  int
}

func ListProfiles() {
	lists := &[]Profile{
		Profile{
			Name: "MAX",
			Age:  19,
		},
		Profile{
			Name: "SUNNY",
			Age:  19,
		},
	}
	fmt.Println("Customers list")
	for i, k := range *lists {
		fmt.Printf("%05d : % 10s, %d\n", i, k.Name, k.Age)
	}
}

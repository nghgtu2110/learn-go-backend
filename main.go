package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var maps = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
		"Hanoi": Vertex{
			21.0278, 105.8342,
		},
		"Hochiminh City": Vertex{
			10.8231, 106.6297,
		},
	}

	fmt.Println(maps)
	fmt.Println("Hanoi at: ", maps["Hanoi"])
	fmt.Println("Hochiminh city at: ", maps["Hochiminh City"])
}

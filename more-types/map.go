package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	{
		m = make(map[string]Vertex)
		m["Bell Labs"] = Vertex{
			40.68433,
			-74.39967,
		}
		fmt.Println(m["Bell Labs"])

	}

	{
		var m = map[string]Vertex{
			"Bell Labs": Vertex{
				40.68433,
				-74.39967,
			},
			"Google": Vertex{
				37.42202,
				-122.08408,
			},
		}
		fmt.Println(m)
	}

	{
		var m = map[string]Vertex{
			"Bell Labs": {
				40.68433,
				-74.39967,
			},
			"Google": {
				37.42202,
				-122.08408,
			},
		}
		fmt.Println(m)
	}

	{
		// mutating maps
		m := make(map[string]int)
		m["Answer"] = 42
		m["Answer"] = 48
		println("The value: ", m["Answer"])
		{
			v, present := m["Answer"]
			println("The value is ", v, "present? ", present)
		}

		delete(m, "Answer")
		println("The value: ", m["Answer"])

		v, present := m["Answer"]
		println("The value is ", v, "present? ", present)
	}
}

// MODIFY THIS AREA
// [filename]: a2_q2.go
// [question]: 2
// [instagram username]: [username]
package main

type Purchase struct {
	order  []int
	budget float64
}

func main() {
	names := []string{"Chips", "Coca Cola", "Chocolate", "Gum", "Water", "Pepsi", "Monster", "Sprite"}
	prices := []float64{2.0, 2.25, 1.75, 1.5, 2.25, 1.75, 3.75, 2.25}

	p1 := Purchase{
		budget: 2.25,
		order:  []int{1},
	}
	p2 := Purchase{
		budget: 8.68,
		order:  []int{7, 4, 1, 6},
	}
	p3 := Purchase{
		budget: 13.20,
		order:  []int{3, 1, 5, 7, 6},
	}
}

func Buy() {
	// Uncomment these lines and use it in your Buy function.
	// Position of these lines vary on your code.

	// answer, _ := reader.ReadString('\n')
	// answer = strings.TrimSuffix(answer, "\n") // Returns "Y" or "N"
}

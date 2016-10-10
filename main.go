package main

import "fmt"

func main() {
	var (
		plantCapacities         = []float64{30, 30, 30, 60, 60, 100}
		activePlants            = []int{0, 1}
		gridLoad        float64 = 75.
	)

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please select an option: ")

	var option string

	fmt.Scanln(&option)

	switch option {
	case "1":
		for index, capacity := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", index, capacity)
		}

	case "2":
		capacity := 0.
		for _, plantId := range activePlants {
			capacity += plantCapacities[plantId]
		}

		utilization := gridLoad / capacity
		fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
		fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
		fmt.Printf("%-20s%.1f%%\n", "Utilization:", utilization*100)

	default:
		fmt.Println("Unsupported option.")
	}

}

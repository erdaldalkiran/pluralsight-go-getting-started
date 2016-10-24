package main

import (
	"fmt"
	"strings"
)

func main() {
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavaliable},
		PowerPlant{solar, 40, inactive},
	}

	grid := PowerGrid{300, plants}

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please select an option: ")

	var option string

	fmt.Scanln(&option)

	switch option {
	case "1":
		grid.generatePlantReport()
	case "2":
		grid.generateGridReport()
	default:
		fmt.Println("Unsupported option.")
	}

}

func generatePlantPowerReport(plantCapacities ...float64) {
	for index, capacity := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", index, capacity)
	}
}

func generateGridPowerReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}

	utilization := gridLoad / capacity
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization:", utilization*100)

}

type PlantType string

const (
	hydro PlantType = "hydro"
	wind  PlantType = "wind"
	solar PlantType = "solar"
)

type PlantStatus string

const (
	active      PlantStatus = "active"
	inactive    PlantStatus = "inactive"
	unavaliable PlantStatus = "unavaliable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (powerGrid *PowerGrid) generatePlantReport() {
	for index, plant := range powerGrid.plants {
		label := fmt.Sprintf("%s%d", "Plant #", index)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type:", plant.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity:", plant.capacity)
		fmt.Printf("%-20s%s\n", "Status:", plant.status)
		fmt.Println("")
	}
}

func (powerGrid *PowerGrid) generateGridReport() {
	capacity := 0.
	for _, plant := range powerGrid.plants {
		if plant.status == active {
			capacity += plant.capacity
		}
	}
	utilization := powerGrid.load / capacity

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", powerGrid.load)
	fmt.Printf("%-20s%.1f%%\n", "Utilization:", utilization*100)
}

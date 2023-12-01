package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	/*
		Create a costs (0.0 filled) array of floats with 31 days (1 day for each input possible)
		Loop through costs
			Check for highest day count -> update if approprate
			Update cost
		Return slice of costs[:highestDayCount]
	*/
	totalDailyCosts := [31]float64{}
	lastDay := 0
	for i := 0; i < len(costs); i++ {
		currentDay := costs[i].day
		if currentDay > lastDay {
			lastDay = currentDay
		}
		totalDailyCosts[currentDay] += costs[i].value
	}
	return totalDailyCosts[:lastDay+1]
}

// dont edit below this line

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}

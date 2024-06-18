package main

import (
	"cmp"
	"slices"
)

type Car struct {
	position, speed int
	ETA             float64
}

func sortCar(a, b Car) int {
	return cmp.Compare(b.position, a.position)
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = Car{position: position[i], speed: speed[i], ETA: float64(target-position[i]) / float64(speed[i])}
	}
	slices.SortFunc(cars, sortCar)
	//fmt.Println(cars)
	fleets := 0
	for len(cars) >= 1 {
		top := cars[0]
		cars = cars[1:len(cars)]
		if len(cars) == 0 {
			fleets += 1
		} else if top.ETA >= cars[0].ETA {
			cars[0] = top
		} else {
			fleets += 1
		}
		//fmt.Println(fleets, cars)
	}
	return fleets
}

func main() {
	//target := 12
	//position := []int{10, 8, 0, 5, 3}
	//speed := []int{2, 4, 1, 1, 3}
	//
	//target := 100
	//position := []int{0, 2, 4}
	//speed := []int{4, 2, 1}

	//target := 10
	//position := []int{3}
	//speed := []int{3}

	target := 14
	position := []int{12, 10, 8, 0, 5}
	speed := []int{1, 3, 4, 1, 1}

	println(carFleet(target, position, speed))
}

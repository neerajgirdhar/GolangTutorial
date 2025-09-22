package main

import (
	"fmt"
)

type simpleCar struct {
	make  string
	model int
}

var simplecar = simpleCar{"Audi", 2020}

type nestedCar struct {
	make       string
	model      int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	sizeininch int
	material   string
}

type embededCar struct {
	wheel
	make  string
	model int
}

func isWheelAlloy(wheel wheel) bool {
	if wheel.material == "Alloy" {
		return true
	}
	return false
}

func (wheel wheel) isWheelAlloy1() bool {
	if wheel.material == "Alloy" {
		return true
	}
	return false
}

func main() {
	// anonymous struct when you have to just create one instance of strut
	//:= works in at function level not at package level
	mycar := struct {
		make  string
		model int
	}{make: "BMW", model: 2022}

	frontwheel := wheel{19, "Alloy"}
	backwheel := wheel{19, "Wrought Iron"}
	nestedcar := nestedCar{"BMW", 2022, frontwheel, backwheel}
	embededcar := embededCar{frontwheel, "Merc", 2023}

	fmt.Println()
	fmt.Println("Simple Strut example ")
	fmt.Printf("Simple Car Model %d\n", simplecar.model)
	fmt.Printf("Simple Car Make %d\n", simplecar.make)

	fmt.Println()
	fmt.Println("Nested Strut example ")
	fmt.Printf("Nested Car Model %d\n", nestedcar.model)
	fmt.Printf("Nested Car Make %d\n", nestedcar.make)
	fmt.Printf("Nested Car Model %d\n", nestedcar.frontWheel.sizeininch)
	fmt.Printf("Nested Car Model %d\n", nestedcar.frontWheel.material)
	fmt.Printf("Nested Car Model %d\n", nestedcar.backWheel.sizeininch)
	fmt.Printf("Nested Car Model %d\n", nestedcar.backWheel.material)
	fmt.Printf("Nested Car Front Wheel Alloys %v\n", isWheelAlloy(nestedcar.frontWheel))
	fmt.Printf("Nested Car Back Wheel Alloys %v\n", isWheelAlloy(nestedcar.backWheel))

	fmt.Println()
	fmt.Println("Embeded Strut example ")
	fmt.Printf("Embeded Car Model %d\n", embededcar.model)
	fmt.Printf("Embeded Car Make %s\n", embededcar.make)
	fmt.Printf("Embeded Car Model %d\n", embededcar.sizeininch)
	fmt.Printf("Embeded Car Model %s\n", embededcar.material)

	fmt.Println()
	fmt.Println("anonymous Strut example ")
	fmt.Printf("anonymous Car Model %d\n", mycar.model)
	fmt.Printf("anonymous Car Make %s\n", mycar.make)

}

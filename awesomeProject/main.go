package main

import (
	"fmt"
)

func multiplyFunctionOutcomeWithTheValue(x1 int, savefuncinvar int) int {
	return x1 * savefuncinvar

}

func multiplereturn(number1, number2 int) (multipliedvalue, addedvalue int) {
	multipliedvalue = number1 * number2
	addedvalue = number2 + number1
	//one Way
	//return multipliedvalue, addedvalue
	// 2nd way
	return
}

func main() {
	fmt.Println("Welcome to Go playground")

	fmt.Println()
	fmt.Println("save Function as a variable and then call that function from variable ")
	savefuncinvar := func(x int, y int) int {
		return x * y
	}
	fmt.Printf("The product is %d\n", savefuncinvar(2, 3))

	fmt.Println()
	fmt.Println("save Function as a variable and pass the parameters from runtime then call that function from variable ")
	savefuncinvar1 := func(x int, y int) int {
		return x * y
	}(3, 3)
	fmt.Printf("The product is %d\n", savefuncinvar1)

	fmt.Println()
	fmt.Println("save Function as a variable and pass the variable/function into another function ")
	doubleValue := multiplyFunctionOutcomeWithTheValue(2, savefuncinvar(2, 3))
	fmt.Printf("The product is %d\n", doubleValue)

	tripleValue := multiplyFunctionOutcomeWithTheValue(3, savefuncinvar1)
	fmt.Printf("The product is %d\n", tripleValue)

	thriceAndNegateValue := multiplyFunctionOutcomeWithTheValue(-3, savefuncinvar1)
	fmt.Printf("The product is %d\n", thriceAndNegateValue)

	fmt.Println()
	fmt.Println("Define multiple variable in same line ")
	fname, lname, age := "Neeraj", "Girdhar", 12
	fmt.Println(fname, lname, age)

	fmt.Println()
	fmt.Println("In Go you can print the type of a variable also ")
	fmt.Printf("Type of fname variable %T\n", fname)
	fmt.Printf("Type of age variable %T\n", age)

	fmt.Println()
	fmt.Println("Multiple return from function ")
	multiply, added := multiplereturn(2, 3)
	fmt.Printf("product %d addition %d \n", multiply, added)

	fmt.Println()
	fmt.Println("Multiple return from function but ignore one return type with underscore as go don't allow unused variables ")
	multiply1, _ := multiplereturn(2, 3)
	fmt.Printf("product %d\n", multiply1)

}

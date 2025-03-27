package main

import (
	"fmt"
	"pujan/day2/auth"
)

func main(){
	// sum, length, success := add(19,10, 5, 13)
	// fmt.Println("The sum is : ",sum)
	// fmt.Println("The length is : ", length)
	// fmt.Println(success)

	// Anonymous & closure function
	// - closure : capture and return
	// func (){
	// 	fmt.Println("I am anonymous")
	// }()

	// display := func(name string){
	// 	fmt.Printf("Name : %s\n", name)
	// }

	// display("Pujan")

	// next := sequenceGenerator()
	// a := next()
	// fmt.Println(a)
	// b := next()
	// fmt.Println(b)
	// for i:=0; i<10; i++{
	// 	fmt.Println(next())
	// }

	// cres := calculate("+", 1,2,3)
	// fmt.Println(cres)

	msg := auth.Login("Pujan")
	fmt.Println(msg)
	
	
}

// variadic function
// func add(numbers ...int) (int,int, string) {
// 	fmt.Println(numbers)
// 	sum := 0
// 	for _, number := range numbers{
// 		sum += number
// 	}
// 	return sum, len(numbers), "added successfully"
// }

// func sequenceGenerator() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }


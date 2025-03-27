package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello World")

	// var a = 10
	// log.Println("Value of a :")
	// log.Printf("a = %d", a)

	// b := "Apple"
	// log.Printf("b = %s", b)
	// b = "Ball"
	// log.Printf("b = %s", b)

	// check := false
	// if check {
	// 	fmt.Println("Checked")
	// } else {
	// 	fmt.Println("UnChecked")
	// }

	// c := 70
	// if c < 50 {
	// 	fmt.Printf("%d is less than 50", c)
	// } else if c > 50 {
	// 	fmt.Printf("%d is greater than 50", c)
	// } else {
	// 	fmt.Printf("%d is 0", c)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Index = %d\n", i+1)
	// }

	// numbers := []int{1, 5, 6, 9}
	// for index, number := range numbers {
	// 	result := fmt.Sprintf("index = %d, number = %d\n", index, number)
	// 	log.Printf("Result = %s", result)
	// }

	if result, value := checkEven(5); result {
		fmt.Printf("Value = %s", value)
	} else {
		fmt.Printf("Value = %s", value)
	}

}

// single return
//	func checkEven(a int) bool {
//		if a%2 == 0 {
//			return true
//		} else {
//			return false
//		}
//	}
//
// multiple return
func checkEven(a int) (bool, string) {
	if a%2 == 0 {
		return true, "Even"
	} else {
		return false, "Odd"
	}
}

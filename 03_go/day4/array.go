package day4

import "fmt"

func StudyArray(){
	// this is array because it has fixed length i.e. [5]
	arr := [5]int{2,3,4,5,6}
	fmt.Println(arr)
	fmt.Println("Length", len(arr))

	// array is value type, so when assigned to new variable, copy it
	arr1 := arr
	fmt.Println("arr1 = ", arr1)
	arr1[2] = 40
	fmt.Println("Original arr = ", arr)
	fmt.Println("New arr1 = ", arr1, "capacity = ", cap(arr1)) // cap is capacity. it is inbuild
	
}
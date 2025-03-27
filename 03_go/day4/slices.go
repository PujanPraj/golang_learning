package day4

import "fmt"

func StudySlice(){
	// this is slice because, it doesn't has fixed length. Its growable unlike array.
	slc := []int{2,3,4,5,6}
	fmt.Println("slc = ", slc, "length = ", len(slc))
	
	// slice is referencetype, so when assigned to new variable, it only takes reference
	slc1 := slc
	fmt.Println("arr1 = ", slc1)
	slc1[2] = 40
	fmt.Println("original slc = ", slc)
	fmt.Println("New slc = ", slc1, "Capacity = ", cap(slc1))


	// empty slice
	copiedSlc := make([]int, len(slc))
	copy(copiedSlc, slc)
	fmt.Println("\n\nCopied slc = ", copiedSlc)
	copiedSlc[0] = 99
	fmt.Println("Original = ", slc)
	fmt.Println("New copied slc = ", copiedSlc)

	// make function, syntax: make(type, length, capacity)
	slcnew := make([]string, 3, 5)
	fmt.Println("\n\nslc new = ", slcnew, "length = ", len(slcnew), "capacity = ", cap(slcnew))
	slcnew = append(slcnew, "40", "hello", "hi")
	fmt.Println("slc after append : ",slcnew, "length = ", len(slcnew), "capacity = ", cap(slcnew))

}
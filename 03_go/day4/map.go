package day4

import "fmt"

func StudyMap(){
	// json/map
	// key : value pair

	person := map[string]string{
		"name" : "Pujan",
		"college" : "Samriddhi",
	}
	
	fmt.Println("Person info = ", person, "length = ", len(person))

	name := person["name"]
	fmt.Println("\nName  = ", name)
	
	for key, value := range person{
		fmt.Println("Key = ", key, "Value = ", value)
	}

	student := make(map[string]any, 0)
	fmt.Println("\nLength of studetn = ", len(student))
	student["name"] = "Pujan"
	student["age"] = 21
	student["status"] = true
	fmt.Println("Student = ", student)

	// deleting key-value pair
	delete(student, "status")
	fmt.Println("After delete = ", student)

	if v, ok := student["status"];ok {
		fmt.Println("Status exist : ", v)
	} else {
		fmt.Println("Status not exists")
	}

}
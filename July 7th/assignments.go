// 1)Create a map where values must be a structure type

package main
import (
	"fmt"
	
)

// func main(){
// 	mpofmp:=map[int]map[string]string{
// 		1: {
// 			"A": "Apple",
// 			"B": "Banana",
// 		}, 
// 		2:{
// 			"Z": "Zebra", 
// 			"Y": "Yalk",
// 		}, 
// 	}
// 	fmt.Println(mpofmp)
// }		



type Person struct {
	Name string
	address string
	fav_number int
}

func main(){
	person1 := Person{
		Name: "James",
		address: "New York",
		fav_number: 20,
	}
	person2 := Person{
		Name: "Mike",
		address: "Chicago",
		fav_number: 3,
	}


//making a map
people := map[string]Person{
	"one": person1,
	"two": person2,
}

fmt.Println(people)



}

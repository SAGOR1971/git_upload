package main

import (
	"fmt"
)

func main() {
	// // simple switch
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday(){
	// case time.Friday, time.Saturday:
	// 	fmt.Println("It's weekend")
	// default:
	// 	fmt.Println("It's workday")
	// }

	// type switch

	whoami := func (i interface{})  {
		switch i.(type){
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a String")
		case bool:
			fmt.Println("It's a boolen")
		}
		
	}

	whoami(true)


}
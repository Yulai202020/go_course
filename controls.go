package main

import (
	"fmt"
)

// if exeption {
	// if body
// }else{
	// else body
// }


func main(){
	var i int
	fmt.Println("Input number 0-10")
	fmt.Scan(&i)
	if i == 7 {
		fmt.Println("It's my favourite number !!!")
	}
	if i > 10 {
		fmt.Println("It's number great than 10 ...")
	}
	if i < 0 {
		fmt.Println("It's number less than 0 ...")
	}

}
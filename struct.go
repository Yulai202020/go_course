package main

import (
	"fmt"
)

type Student struct{
	name string
	surname string
	class int
}

func main(){
	Yulai := Student{
		name:"Yulai",
		surname:"Nigmatullin",
		class:6,
	}

	fmt.Println(Yulai.name)
	fmt.Println(Yulai.surname)
	fmt.Println(Yulai.class)
}
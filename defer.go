package main

import (
	"fmt"
)

// defer - говорит запустит строку в конце

func main(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
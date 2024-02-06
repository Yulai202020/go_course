package main

import (
	"fmt"
)

// const ConstName = value
// var name1 type = value // or
// var name2 = value

// nil = None (null)

func main(){
	const MyConst  = 10
	var hello string = "Hello world!!!"
	var equaltion1 bool = 1 == 1
	var equaltion2 bool = 1 == 2

	var arr1 = [10]int{98, 99, 100} // array [колво елементов]тип{данные}
	var arr2 = []int{98, 99, 100,1212} // array
	fmt.Println(arr1)
	fmt.Println(arr2)

	dict := make(map[string]int) // словарь map[key_type]
	dict = map[string]int{
		"Moscow": 9000000,
	}
	fmt.Println(len(dict))
}
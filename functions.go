package main

import (
	"fmt"
)

func sum(values ...int) (result int) {
	result = 0
	for _, v := range values {
		result += v
	}
	return result
}

func main(){
	fmt.Println(sum(1,2,3,4,5))
}
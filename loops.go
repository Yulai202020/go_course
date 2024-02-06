package main

import (
	"fmt"
)

func main(){
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	s := "hello GO !1!"
	for k,v := range s {
		fmt.Println(k,string(v))
	}
}
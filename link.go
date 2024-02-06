package main

import (
	"fmt"
)

func main(){
	var a int = 49
	var b *int = &a // ссылка на a
	fmt.Println(a, *b) // b - это ссылка чтобы взять содержимое надо писать *b

	a = 25
	fmt.Println(a, *b) // изменяем a и смотрим что получется

	*b = 16
	fmt.Println(a, *b) // изменяем содержимое b и смотрим что получется

}
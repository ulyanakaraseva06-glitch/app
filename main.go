package main

import (
	"app-1/feature"
	"fmt"
)


func main() {
	num := 10
	text := "lox"
	num++
	text2 := " bugaga"
	text += text2
	fmt.Println(num, text)
	bub(5, 5)
	feature.Srez()
}


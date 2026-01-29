package main

import "fmt"

func main() {
	num := 10
	text := "lox"
	num++
	text2 := " bugaga"
	text += text2
	fmt.Println(num, text)
}

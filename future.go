package main

import (
	"app-1/feature"
	"fmt"
)

func bub(a, b int) {
	max, err := feature.New_f(a, b)
	if err != nil {
		err = fmt.Errorf("Ошибка: %w", err)
		fmt.Println(err)
	}
	fmt.Printf("Максимальное число: %v", max)
}

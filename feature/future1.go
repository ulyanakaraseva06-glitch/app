package feature

import "fmt"

func New_f(a int, b int) (int, error) {
	switch {
	case a == b:
		return 0, fmt.Errorf("number are equal")
	default:
		fmt.Println("переменные не равны")
	}

	if a > b {
		return a, nil
	}
	return b, nil

}

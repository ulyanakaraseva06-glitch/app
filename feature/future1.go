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

// 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000
// 100, 200, 300, 500, 600, 700, 800, 900, 1000, 1000(так как в базовом массиве на крайнем индексе была 1000,
//после удаления 400 она остается на своем месте)
func Srez() {
	prices := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	fmt.Println(prices[0:5])//тут кароче первые пять цифр выводим
	fmt.Println(prices[len(prices)-3:])//а тут три последние
	fmt.Println(prices[1:8])//тут выводим с 2 по 7 число включительно
	price1 := append(prices[0:3], prices[4:]...)
	fmt.Println(price1)//удаляем 400
	prices = append(prices, 1100, 1200)
	fmt.Println(prices)//добавляем 1100, 1200(изменяем базовый массив)
}

package feature

import (
	"fmt"
)

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
	price := [8]int{1, 2, 3,4, 5, 6, 7, 8}
	fmt.Println(price[0:1])
	fmt.Println(price[0:0])
	Cickl()
}

func Cickl(){
	for i := 1; i < 10; i++{
		fmt.Println(i*i)//здесь мы выводим квадраты чисел
	}
	strw := "Marina"

	for index, value := range strw{
		fmt.Println("index", index, "value", value)
	}
	for _, value := range strw{
		fmt.Println("value", value)
	}
	for index, _ := range strw{
		fmt.Println("index", index)
	}
	var i = 1
	for i <= 5 {
		fmt.Println(i*2)
		i++
	}
	im := 1
	for {
		fmt.Println(im+1)
		im++
		if im >= 5 {
			break
		}
	}
	b := true
	if !b{
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
	spisok := make([]int, 100)
	for iq := 1; iq <= 100; iq++{
		spisok[iq-1]=iq

	}
	fmt.Print(spisok)


}
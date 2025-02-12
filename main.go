package main

import (
	"fmt"
	"strconv"
)

//Калькулятор

// func main() {
// 	var num1, num2 int
// 	var do string
// 	for {
// 		fmt.Print("Напишите первое число: ")
// 		fmt.Scan(&num1)
// 		fmt.Print("Напишите второе число: ")
// 		fmt.Scan(&num2)
// 		fmt.Print("Введите действие(+, -, /, *): ")
// 		fmt.Scan(&do)
// 		switch {
// 		case do == "+":
// 			plus(num1, num2)
// 		case do == "-":
// 			minus(num1, num2)
// 		case do == "*":
// 			umn(num1, num2)
// 		case do == "/":
// 			del(num1, num2)
// 		}
// 		isRepeatCalculetion := checkrepeatCalculation()
// 		if !isRepeatCalculetion {
// 			break
// 		}
// 	}

// }

// func plus(num1 int, num2 int) {
// 	res := num1 + num2
// 	fmt.Print(res)
// }

// func minus(num1 int, num2 int) {
// 	res := num1 - num2
// 	fmt.Print(res)
// }

// func del(num1 int, num2 int) {
// 	res := num1 / num2
// 	fmt.Print(res)
// }

// func umn(num1 int, num2 int) {
// 	res := num1 * num2
// 	fmt.Print(res)
// }

// func checkrepeatCalculation() bool {
// 	var userChoise string
// 	fmt.Println("Вы хотите сделать еще расчет? (Y/n)")
// 	fmt.Scan(&userChoise)

// 	if userChoise == "y" || userChoise == "Y" {
// 		return true
// 	}
// 	return false
// }

func main() {
	str := []string{"10", "10.5", "sfsdfds", "345"}
	res := summNumber(str)
	fmt.Println("Сумма чисел: ", res)
}

func summNumber(str []string) float64 {
	var summ float64
	for _, s := range str {
		num, _ := strconv.ParseFloat(s, 64) // Игнорируем ошибку, она всегда будет игнорироваться
		summ += num
	}
	return summ
}

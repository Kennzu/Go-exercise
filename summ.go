package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := []string{"10", "10.5", "sfsdfds", "345"}
	res := summNumber(str)
	fmt.Println("Сумма чисел: ", res)
}

func summNumber(str []string) float64 {
	var summ float64
	for _, s := range str {
		num, _ := strconv.ParseFloat(s, 64)
		summ += num
	}
	return summ
}

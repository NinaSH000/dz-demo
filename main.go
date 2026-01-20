package main

import "fmt"

func main() {
	const USDtoEUR float64 = 0.25
	const USDtoRUB float64 = 84
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
	sum, currencyFrom, currencyTo := inputData()
	fmt.Println(sum, currencyFrom, currencyTo)
}

func inputData() (float64, string, string) {
	var sum float64
	var currencyFrom string
	var currencyTo string

	fmt.Println("сколько денег конвертировать: ")
	fmt.Scan(&sum)
	fmt.Println("Валюта 1: ")
	fmt.Scan(&currencyFrom)
	fmt.Println("Валюта 2: ")
	fmt.Scan(&currencyTo)
	return sum, currencyFrom, currencyTo
}

func calculation(sum float64, currencyFrom string, currencyTo string) {

}

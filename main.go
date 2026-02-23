package main

import (
	"fmt"
	"errors"
)

func main() {
	sum, currencyFrom, currencyTo := inputData()
	fmt.Println(sum, currencyFrom, currencyTo)
	money, err := calculation(sum, currencyFrom, currencyTo)
	if err != nil {
			fmt.Println(err)
		}
	fmt.Println(money)
}

func inputData() (float64, string, string) {
	meter := 1

	var sum float64
	var currencyFrom string
	var currencyTo string

	for {
		if meter == 1 {
			fmt.Println("сколько денег конвертировать: ")
			fmt.Scan(&sum)

			if sum > 0 {
				meter++
			} else {
				fmt.Println("Введите заново")
				continue
			}
		}
		if meter == 2 {
			fmt.Println("Валюта 1(EUR, RUB, USD): ")
			fmt.Scan(&currencyFrom)
			if currencyFrom == "RUB" || currencyFrom == "rub" || currencyFrom == "EUR" || currencyFrom == "eur" || currencyFrom == "USD" || currencyFrom == "usd" {
				meter++
			} else {
				fmt.Println("Введите заново")
				continue
			}
		}

		if meter == 3 {
			if currencyFrom == "RUB" || currencyFrom=="rub"{
				fmt.Println("Валюта 2 (EUR, USD): ")
				fmt.Scan(&currencyTo)
				if currencyTo == "EUR" || currencyTo == "eur" || currencyTo == "USD" || currencyTo == "usd" {
					return sum, currencyFrom, currencyTo
				} else {
					fmt.Println("Введите заново")
					continue
				}
			}
			if currencyFrom == "EUR" || currencyFrom=="eur"{
				fmt.Println("Валюта 2 (RUB, USD): ")
				fmt.Scan(&currencyTo)
				if currencyTo == "RUB" || currencyTo == "rub" || currencyTo == "USD" || currencyTo == "usd" {
					return sum, currencyFrom, currencyTo
				} else {
					fmt.Println("Введите заново")
					continue
				}
			}
			if currencyFrom == "USD" || currencyFrom=="usd"{
				fmt.Println("Валюта 2 (RUB, EUR): ")
				fmt.Scan(&currencyTo)
				if currencyTo == "RUB" || currencyTo == "rub" || currencyTo == "EUR" || currencyTo == "eur" {
					return sum, currencyFrom, currencyTo
				} else {
					fmt.Println("Введите заново")
					continue
				}
			}
		}
	}
}

func calculation(sum float64, currencyFrom string, currencyTo string) (float64, error) {
	const USDtoEUR float64 = 1.18
	EURtoUSD := 1/USDtoEUR
	const USDtoRUB float64 = 84
	RUBtoUSD  := 1/USDtoRUB
	EURtoRUB := USDtoRUB / USDtoEUR
	RUBtoEUR := 1/EURtoRUB

	if currencyFrom == "EUR" || currencyFrom == "eur" {
		if currencyTo == "USD" || currencyTo == "usd"{
			money := sum*EURtoUSD
			return money, nil
		}
		if  currencyTo == "RUB" || currencyTo == "rub"{
			money := sum*EURtoRUB
			return money, nil
		}
	}

	if currencyFrom == "RUB" || currencyFrom == "rub" {
		if currencyTo == "USD" || currencyTo == "usd"{
			money := sum*RUBtoUSD
			return money, nil
		}
		if  currencyTo == "EUR" || currencyTo == "eur"{
			money := sum*RUBtoEUR
			return money, nil
		}
	}

	if currencyFrom == "USD" || currencyFrom == "usd" {
		if currencyTo == "EUR" || currencyTo == "eur"{
			money := sum*USDtoEUR
			return money, nil
		}
		if  currencyTo == "RUB" || currencyTo == "rub"{
			money := sum*USDtoRUB
			return money, nil
		}
	}
	return 0.0 , errors.New("Что-то пошло не так")
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"os"
	"strconv"
    "strings"
)


func main() {
	
	for {
	fmt.Println("что посчитать? среднее(1), сумма(2), медиана(3):")
	answer := 0
	fmt.Scan(&answer)
		if answer == 1{
			transaction := input()
			avg := getAvg(transaction)
			fmt.Printf("%.2f\n", avg)
			break
		}
		if answer == 2{
			transaction := input()
			sum := getSum(transaction)
			fmt.Println(sum)
			break
		}
		if answer == 3{
			transaction := input()
			med := getMed(transaction)
			fmt.Println(med)
			break
		} else {
			continue
		}
	}
}

func getAvg (trans []float64) float64{
	avg := 0.0
	for _, value := range trans {
		avg += value 
	}
	return avg/float64(len(trans))
} //считаем среднее

func getSum (trans []float64) float64{
	sum := 0.0
	for _, value := range trans {
		sum += value
	}
	return sum
} // считаем сумму

func getMed(trans []float64) float64{
	med := 0.0
	sort.Float64s(trans)

	if len(trans)%2==0 {
		a := len(trans)/2
		med = (trans[a]+trans[a-1])/2	
	} else {
		a := int(math.Ceil(float64(len(trans))/2))
		med = trans[a-1]
	}
	return med
} //считаем медиану

func input() []float64 {
	// Создаём reader для ввода с клавиатуры
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите числа через запятую (например: 70, 54, f, 6.9, 0, 52, 1, 1.7, g, 4f):")

	// Читаем всю строку, которую ввёл пользователь
    input, _ := reader.ReadString('\n')

	// Убираем символ переноса строки в конце
    input = strings.TrimSpace(input)

	 // Разделяем по запятой
    parts := strings.Split(input, ",")

	 // Слайс для результатов
    var numbers []float64

	 // Обрабатываем каждую часть
    for _, part := range parts {
        // Убираем пробелы
        clean := strings.TrimSpace(part)
        
        // Пробуем преобразовать в число
        if num, err := strconv.ParseFloat(clean, 64); err == nil {
            numbers = append(numbers, num)
		}
	}
	return numbers
}

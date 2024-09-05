//КАЛЬКУЛЯТОР

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var desert = make(map[string]int)
	desert["I"] = 1
	desert["II"] = 2
	desert["III"] = 3
	desert["IV"] = 4
	desert["V"] = 5
	desert["VI"] = 6
	desert["VII"] = 7
	desert["VIII"] = 8
	desert["IX"] = 9
	desert["X"] = 10
	desert["XI"] = 11
	desert["XII"] = 12
	desert["XIII"] = 13
	desert["XIV"] = 14
	desert["XV"] = 15
	desert["XVI"] = 16
	desert["XVII"] = 17
	desert["XVIII"] = 18
	desert["IXX"] = 19
	desert["XX"] = 20

Start:
	fmt.Println("Введите выражение (через пробел) с помощью арабских или римских чисел (от 1 до 10) с операцией(+, -, /, *) ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text) // Удаляем лишние пробелы и символы новой строки
	slice := strings.Split(text, " ")

	if len(slice) != 3 {
		fmt.Println("Неверный ввод")
		fmt.Println()
		goto Start
	}

	firstelem := slice[0]
	xi := 0
	secondelem := slice[2]
	yi := 0

	flag := false
	if val, ok := desert[firstelem]; ok {
		flag = true
		xi = val
	}
	if val, ok := desert[secondelem]; ok {
		yi = val
	} else {
		flag = false
	}

	if !flag {
		x, err1 := strconv.Atoi(slice[0])
		if err1 != nil {
			log.Fatal("Ошибка преобразования первого числа:", err1)
		}
		y, err2 := strconv.Atoi(slice[2])
		if err2 != nil {
			log.Fatal("Ошибка преобразования второго числа:", err2)
		}
		if x < 0 || x > 10 || y < 0 || y > 10 {
			fmt.Println("Неправельный ввод")
			fmt.Println()
			goto Start
		}
		switch slice[1] {
		case "+":
			fmt.Println(x + y)
		case "-":
			if x >= y {
				fmt.Println(x - y)
			}
		case "*":
			fmt.Println(x * y)
		case "/":
			if y == 0 {
				fmt.Println("Ошибка: деление на ноль.")
			} else {
				fmt.Println(float64(x) / float64(y))
			}
		default:
			fmt.Println("Неизвестный оператор")
		}
	} else {
		result := 0
		switch slice[1] {
		case "+":
			result = xi + yi
		case "-":
			if xi >= yi {
				result = xi - yi
			} else {
				fmt.Println("В данном калькуляторе не предусмотрены отрицательные числа")
			}
		case "*":
			result = xi * yi
		case "/":
			if yi == 0 {
				fmt.Println("Ошибка: деление на ноль.")
			} else {
				result = xi / yi
			}
		default:
			fmt.Println("Неизвестный оператор")
		}
		for key, val := range desert {
			if val == result {
				fmt.Println(key)
				break
			}
		}
	}
}

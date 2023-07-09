package main

import (
	"fmt"
	"strconv"
	"strings"
)

func plus(oper string) string {
	var ch [2]int
	numbers := strings.Split(oper, "+")
	for idx, number := range numbers {
		//	fmt.Printf("chislo%d = %s\n", idx, number)
		ch[idx], err := strconv.Atoi(number)
		if err != nil {
			//о нет что-то пошло не так
		}
	}

	return oper
}
func minus(oper string) string {
	fmt.Print(strings.Split(oper, "-"))
	return oper
}

func umn(oper string) string {
	fmt.Print(strings.Split(oper, "*"))
	return oper
}

func del(oper string) string {
	fmt.Print(strings.Split(oper, "/"))
	return oper
}

func main() {
	fmt.Println("Введите операцию:")
	var a string
	fmt.Scanf("%s", &a)
	fmt.Printf("Тип: %T Значение: %s\n", a, a)
	for i := range a {
		switch a[i] {
		case '+':
			fmt.Println(plus(a))
		case '-':
			fmt.Println(minus(a))
		case '*':
			fmt.Println(del(a))
		case '/':
			fmt.Println(del(a))
		default:
			continue
		}
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func check(num string) int { //Проверка числа
	R := [...]string{"I", "V", "X", "L", "C"}
	A := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for k := range R {
		if num[0:1] == R[k] {
			return 2 //Число римское
		}
	}
	for k := range A {
		if num[0:1] == A[k] {
			return 1 //Число арабское
		}
	}
	uslov()
	return 0
}

func rome(Rim string) int { //Перевод числа на арабский
	chRim := map[string]int{"I": 1, "V": 5, "X": 10}
	chArab := 0
	for i := range Rim {
		if i < len(Rim)-1 && chRim[Rim[i:i+1]] < chRim[Rim[i+1:i+2]] {
			chArab -= chRim[Rim[i:i+1]]
		} else {
			chArab += chRim[Rim[i:i+1]]
		}
	}
	return chArab
}

func arab(Ara int) string { //Перевод числа на римский
	Rim := map[int]string{100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	s := [...]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	ChRim := ""
	i := 0
	for i < 9 {
		if Ara-s[i] >= 0 {
			Ara -= s[i]
			ChRim += Rim[s[i]]
		} else {
			i++
		}
	}
	return ChRim
}

func uslov() string {
	return fmt.Sprintf("Введенные вами числа не соответсвуют условиям программы.")
}

func plus(oper string) string {
	var ch [2]int
	var vd [2]int
	numbers := strings.Split(oper, "+")
	for idx, number := range numbers {
		switch check(number) {
		case 1:
			s, err := strconv.Atoi(number)
			if err != nil {
				//пездец
			}
			ch[idx] = s
			vd[idx] = check(number)
		case 2:
			ch[idx] = rome(number)
			vd[idx] = check(number)
		}
	}
	if ch[0] > 10 || ch[0] < 1 || ch[1] > 10 || ch[1] < 1 || vd[0] != vd[1] {
		return uslov()
	}
	if vd[0] == 1 {
		return fmt.Sprintf("Сумма %d и %d равна %d.", ch[0], ch[1], ch[0]+ch[1])
	} else {
		return fmt.Sprintf("Сумма %s и %s равна %s.", arab(ch[0]), arab(ch[1]), arab(ch[0]+ch[1]))
	}
}
func minus(oper string) string {
	var ch [2]int
	var vd [2]int
	numbers := strings.Split(oper, "-")
	for idx, number := range numbers {
		switch check(number) {
		case 1:
			s, err := strconv.Atoi(number)
			if err != nil {
				//пездец
			}
			ch[idx] = s
			vd[idx] = check(number)
		case 2:
			ch[idx] = rome(number)
			vd[idx] = check(number)
		}
	}
	if ch[0] > 10 || ch[0] < 1 || ch[1] > 10 || ch[1] < 1 || vd[0] != vd[1] {
		return uslov()
	}
	if vd[0] == 1 {
		return fmt.Sprintf("Разность %d и %d равна %d.", ch[0], ch[1], ch[0]-ch[1])
	} else {
		if ch[0] < ch[1] || ch[0] == ch[1] {
			return uslov()
		}
		return fmt.Sprintf("Разность %s и %s равна %s.", arab(ch[0]), arab(ch[1]), arab(ch[0]-ch[1]))
	}
}

func umn(oper string) string {
	var ch [2]int
	var vd [2]int
	numbers := strings.Split(oper, "*")
	for idx, number := range numbers {
		switch check(number) {
		case 1:
			s, err := strconv.Atoi(number)
			if err != nil {
				//пездец
			}
			ch[idx] = s
			vd[idx] = check(number)
		case 2:
			ch[idx] = rome(number)
			vd[idx] = check(number)
		}
	}
	if ch[0] > 10 || ch[0] < 1 || ch[1] > 10 || ch[1] < 1 || vd[0] != vd[1] {
		return uslov()
	}
	if vd[0] == 1 {
		return fmt.Sprintf("Произведение %d и %d равна %d.", ch[0], ch[1], ch[0]*ch[1])
	} else {
		return fmt.Sprintf("Произведение %s и %s равна %s.", arab(ch[0]), arab(ch[1]), arab(ch[0]*ch[1]))
	}
}

func del(oper string) string {
	var ch [2]int
	var vd [2]int
	numbers := strings.Split(oper, "/")
	for idx, number := range numbers {
		switch check(number) {
		case 1:
			s, err := strconv.Atoi(number)
			if err != nil {
				//пездец
			}
			ch[idx] = s
			vd[idx] = check(number)
		case 2:
			ch[idx] = rome(number)
			vd[idx] = check(number)
		}
	}
	if ch[0] > 10 || ch[0] < 1 || ch[1] > 10 || ch[1] < 1 || vd[0] != vd[1] || ch[1] == 0 {
		return uslov()
	}
	if vd[0] == 1 {
		return fmt.Sprintf("Частное %d и %d равна %d.", ch[0], ch[1], ch[0]/ch[1])
	} else {
		return fmt.Sprintf("Частное %s и %s равна %s.", arab(ch[0]), arab(ch[1]), arab(ch[0]/ch[1]))
	}
}

func main() {
	fmt.Println("Введите операцию:")
	var a string
	fmt.Scanf("%s", &a)
	if a[0] == '-' || check(a) == 0 {
		print(uslov())
		return
	}
	for i := range a {
		switch a[i] {
		case '+':
			fmt.Println(plus(a))
		case '-':
			fmt.Println(minus(a))
		case '*':
			fmt.Println(umn(a))
		case '/':
			fmt.Println(del(a))
		default:
			continue
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RomanLib(flag string, tens int, units int) map[string]int32 {
	//Неужели в Go нельзя задать параметр функции с дефолтным значением, что бы можно было не передавать их при вызове ??? К примерру flag string = "default"
	var roman map[string]int32
	roman = make(map[string]int32)
	roman["I"] = 1
	roman["II"] = 2
	roman["III"] = 3
	roman["IV"] = 4
	roman["V"] = 5
	roman["VI"] = 6
	roman["VII"] = 7
	roman["VIII"] = 8
	roman["IX"] = 9
	roman["X"] = 10
	roman["XX"] = 20
	roman["XXX"] = 30
	roman["XL"] = 40
	roman["L"] = 50
	roman["LX"] = 60
	roman["LXX"] = 70
	roman["LXXX"] = 80
	roman["XC"] = 90
	roman["C"] = 100
	var TensOnRoman string
	var UnitsOnRoman string
	if flag == "find" {
		if tens > 0 || units > 0 {
			for k, v := range roman {
				if (tens * 10) == int(v) {
					TensOnRoman = k
				}
				if units == int(v) {
					UnitsOnRoman = k
				}
				//fmt.Printf("key: %d, value: %t\n", k, v)
			}
			if len(TensOnRoman) > 0 {
				if len(UnitsOnRoman) > 0 {
					fmt.Println(TensOnRoman + UnitsOnRoman)
				} else {
					fmt.Println(TensOnRoman)
				}

			} else {
				fmt.Println(UnitsOnRoman)
			}
		} else {
			fmt.Println("Ошибка : Результат Вычесления с римскими операндами не может быть отрицательным")
		}
	}
	return roman
}
func arabicToRoman(val int) {
	var tens = val / 10
	var units = val % 10
	if val > 10 {
		val = val % 10
	}
	// fmt.Println("tens : ", tens, " units : ", units)
	RomanLib("find", tens, units)
}
func calculate(operator string, val1 int, val2 int, flag string) string {
	var result int
	switch operator {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		if val2 != 0 {
			result = val1 / val2
		} else {
			panic("Devision by ZERO ?")
		}
	default:
		panicMsg()
	}
	if flag == "roman" {
		arabicToRoman(result)
	}
	return strconv.Itoa(result)
}
func panicMsg() {
	fmt.Println("!!! Не правильные операнды, либо оператор !!!")
	fmt.Println("______________________________")
	fmt.Println("1) Данные передаются в одну строку")
	fmt.Println("2) Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.")
	fmt.Println("3) Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b")
	fmt.Println("4) Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более.")
	fmt.Println("5) Калькулятор умеет работать только с целыми числами.")
	fmt.Println("6) Калькулятор умеет работать только с арабскими или римскими цифрами одновременно.")
}
func main() {

	roman := RomanLib("def", 0, 0)
	var inputA string
	var operator string
	var inputB string
	var flag string // Флаг режима работы калькулятора
	nr := bufio.NewReader(os.Stdin)
	inputData, _ := nr.ReadString('\n')
	inputDataArr := strings.Split(inputData, " ")
	if len(inputDataArr) == 3 {
		inputA = strings.ToUpper(inputDataArr[0])
		operator = inputDataArr[1]
		inputB = strings.ToUpper(strings.Trim(inputDataArr[2], "\r\n"))
		if strings.Contains(inputA, ".") || strings.Contains(inputA, ",") || strings.Contains(inputB, ".") || strings.Contains(inputB, ",") {
			panic("Калькулятор работает только с целыми числами ! ")
		}
	} else {
		fmt.Println("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	var answer string
	if roman[inputA] > 0 || roman[inputB] > 0 {
		//fmt.Printf("is Roman numbers")
		if (roman[inputA] > 0) && (roman[inputB] > 0) {
			if roman[inputA] > 10 || roman[inputB] > 10 {
				panic("Данные на вход римских от 1 до 10 включительно")
			}
			flag = "roman" // Устанавливаем флаг, что мы работаем в римской системе
			answer = calculate(operator, int(roman[inputA]), int(roman[inputB]), flag)
		} else {
			panic("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно. \r\nОдин аргумент не римское число")
		}
	} else {
		arabA, erA := strconv.Atoi(inputA)
		arabB, erB := strconv.Atoi(inputB)
		if (erA != nil) || (erB != nil) {
			panicMsg()
		}
		if (arabA >= 1 && arabA <= 10) && (arabB >= 1 && arabB <= 10) {
			answer = calculate(operator, arabA, arabB, "null")

		} else {
			panic("Данные на вход от 1 до 10 включительно")
		}

	}
	// Тут выводим финальный ответ, если режим работы калькулятора не римский, т.к ответ римскими выводится в функции где ответ конвертируется в римские
	if flag != "roman" {
		fmt.Println(answer)
	}

}

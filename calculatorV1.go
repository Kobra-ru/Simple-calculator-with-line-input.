package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	/*
		Лист возможных дополнений/улучшений/уточнения:
			1)Не работает с запятыми также в расчётах/ответе не использует запятые()
			2)Ошибка огромных чисел
			3)Малое кол-во операторов
			4)Ошибочная ошибка при вводе оператора **
			5)Константы (число пи и тд.)
			6)Обработка букв насколько нужна большой вопрос?
			7)Оптимизация?Пару if/циклов удалить ну хз

	*/
	en := true
	line := ""
	fmt.Println("!Внимание данный калькулятор не работает с числами с запятой и не использует их в расчётах/ответах, возможна потеря точности при делении.")
	fmt.Printf("Здравствуйте вы попали в калькулятор, чтобы выйти пропишите exit.")
	fmt.Println("Доступные операторы: + - * / ( )")
	for en {
		line = ""
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line = scanner.Text()
		if line == "exit" {
			en = false
		} else {
			var sum int
			var errorCode int
			sum, errorCode = calc(line)
			if errorCode == 0 {
				fmt.Printf("Ответ: %d\n", sum)
			} else {
				switch errorCode {
				case 1:
					fmt.Println("Error: Синтаксическая ошибка")
				case 2:
					fmt.Println("Error: Неподдерживаемый оператор или Неподдерживаемый символ")
				case 3:
					fmt.Println("Error: Нет входных данных")
				case 4:
					fmt.Println("Error: Деление на ноль")
				}
			}

		}
		var test float64
		test = 153463425.5224424242
		fmt.Printf("%0.4f\n", math.Round(test*10000)/10000)
	}
}
func stringToAction(line string) ([]int, int) { // обрабатывает ошибки Cинтаксическая, Неподдерживаемый оператор или Неподдерживаемый символ,
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )

	var arr []int // основной массив
	var temp string
	var oldSym rune
	if line != "" {
		for i := 0; i < len(line); i++ {
			sym := rune(line[i])
			if sym == '+' || sym == '-' || sym == '*' || sym == '/' || sym == '(' || sym == ')' {
				if temp != "" {
					arr = append(arr, stringToNumber(temp))
					temp = ""
				}
				switch sym {
				case '+':
					arr = append(arr, -1)
				case '-':
					arr = append(arr, -2)
				case '*':
					arr = append(arr, -3)
				case '/':
					arr = append(arr, -4)
				case '(':
					arr = append(arr, -5)
				case ')':
					arr = append(arr, -6)
				}

			} else if oldSym == ' ' && isNumber(rune(sym)) == true && temp != "" {
				return arr, 1
			} else if isNumber(rune(sym)) == true {
				temp = temp + string(sym)
			} else if sym != ' ' {
				return arr, 2
			}
			oldSym = rune(line[i])
		}
		if temp != "" {
			arr = append(arr, stringToNumber(temp))
			temp = ""
		}
	}
	if len(arr) == 0 {
		return []int{}, 3
	}
	return arr, 0
}
func calc(line string) (int, int) { // обрабатывает ошибки Cинтаксическая, Деление на ноль
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )
	arr, errorCode := stringToAction(line)
	//fmt.Println("Тест массив:", arr) //test

	if errorCode != 0 {
		return 0, errorCode
	}

	arr, errorCode = recursionOfStaples(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	arr = unaryMinus(arr)

	arr, errorCode = multiplicationOrDivision(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	arr, errorCode = plusOrMinus(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	if len(arr) == 0 || (arr[0] >= -6 && arr[0] < 0) || len(arr) > 1 { //ловля ошибок
		return 0, 1
	}

	if arr[0] < -6 { // 				декодирования минусовых ответов
		arr[0] = arr[0] + 6
	}

	return arr[0], errorCode
}

func calcStaples(arr []int) (int, int) {
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )
	errorCode := 0

	arr, errorCode = recursionOfStaples(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	arr = unaryMinus(arr)

	arr, errorCode = multiplicationOrDivision(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	arr, errorCode = plusOrMinus(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	if len(arr) == 0 || (arr[0] >= -6 && arr[0] < 0) || len(arr) > 1 { //ловля ошибок
		return 0, 1
	}

	return arr[0], 0
}
func calcOper(number1 int, oper int, number2 int) (int, int) {
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )
	if number1 < -6 { //нормалезуем
		number1 = number1 + 6
	}
	if number2 < -6 {
		number2 = number2 + 6
	}
	sum := 0
	switch oper {
	case -1:
		sum = number1 + number2
	case -2:
		sum = number1 - number2
	case -3:
		sum = number1 * number2
	case -4:
		if number2 == 0 {
			return 0, 4
		} else {
			sum = number1 / number2
		}
	}

	if sum < 0 { //кодируем
		sum -= 6
	}
	return sum, 0
}
func isNumber(str rune) bool {
	if str == '0' || str == '1' || str == '2' || str == '3' || str == '4' || str == '5' || str == '6' || str == '7' || str == '8' || str == '9' {
		return true
	}
	return false
}
func stringToNumber(str string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '1':
			result = result + 1*int(math.Pow(10, float64(len(str)-1-i)))
		case '2':
			result = result + 2*int(math.Pow(10, float64(len(str)-1-i)))
		case '3':
			result = result + 3*int(math.Pow(10, float64(len(str)-1-i)))
		case '4':
			result = result + 4*int(math.Pow(10, float64(len(str)-1-i)))
		case '5':
			result = result + 5*int(math.Pow(10, float64(len(str)-1-i)))
		case '6':
			result = result + 6*int(math.Pow(10, float64(len(str)-1-i)))
		case '7':
			result = result + 7*int(math.Pow(10, float64(len(str)-1-i)))
		case '8':
			result = result + 8*int(math.Pow(10, float64(len(str)-1-i)))
		case '9':
			result = result + 9*int(math.Pow(10, float64(len(str)-1-i)))

		}
	}
	return result
}
func recursionOfStaples(arr []int) ([]int, int) {
	var newArr []int
	var tempArr []int
	flag := false
	quantity := 0
	for i := 0; i < len(arr); i++ { // скобки
		errorCode := 0
		if arr[i] == -6 && flag == true {
			quantity--
		} else if arr[i] == -5 && flag == true {
			quantity++
		} else if arr[i] == -5 && flag == false {
			flag = true
			quantity++
		} else if flag == false {
			newArr = append(newArr, arr[i])
		}
		if flag == true && quantity == 0 {
			flag = false
			var temp int
			temp, errorCode = calcStaples(tempArr)
			newArr = append(newArr, temp)
			tempArr = []int{}
		}
		if flag == true && quantity != 0 && !(quantity == 1 && arr[i] == -5) {
			tempArr = append(tempArr, arr[i])
		}
		if errorCode != 0 {
			return []int{}, errorCode
		}
	}
	if len(newArr) == 0 || quantity != 0 { //ловля ошибок
		return []int{}, 1
	}
	return newArr, 0
}
func unaryMinus(arr []int) []int {
	newArr := []int{}
	for i := 0; i < len(arr); i++ { //					унарный -

		if i == 0 && i+1 < len(arr) && arr[i] == -2 && (arr[i+1] >= 0 || arr[i+1] < -6) {
			if arr[i+1] == 0 {
				newArr = append(newArr, 0)
			} else {
				newArr = append(newArr, arr[i+1]*-1-6)
			}
			i++
		} else if i > 0 && i+1 < len(arr) && arr[i] == -2 && arr[i-1] <= -1 && arr[i-1] >= -5 && (arr[i+1] >= 0 || arr[i+1] < -6) {
			if arr[i+1] == 0 {
				newArr = append(newArr, 0)
			} else {
				newArr = append(newArr, arr[i+1]*-1-6)
			}
			i++
		} else {
			newArr = append(newArr, arr[i])
		}

	}
	return newArr
}
func multiplicationOrDivision(arr []int) ([]int, int) {
	newArr := []int{}
	stop := 0
	for i := 0; i < len(arr); i++ { // умножения деления
		errorCode := 0
		stop--
		if i+2 < len(arr) && (arr[i] >= 0 || arr[i] < -6) && (arr[i+1] == -3 || arr[i+1] == -4) && (arr[i+2] >= 0 || arr[i+2] < -6) {
			if len(newArr) > 0 && (newArr[len(newArr)-1] >= 0 || newArr[len(newArr)-1] < -6) && stop == 1 {
				newArr[len(newArr)-1], errorCode = calcOper(newArr[len(newArr)-1], arr[i+1], arr[i+2])
				i++
			} else {
				var temp int
				temp, errorCode = calcOper(arr[i], arr[i+1], arr[i+2])
				newArr = append(newArr, temp)
				i++
			}
			stop = 2
		} else if stop <= 0 {
			newArr = append(newArr, arr[i])
		}
		if errorCode != 0 {
			return []int{}, errorCode
		}

	}
	return newArr, 0
}
func plusOrMinus(arr []int) ([]int, int) {
	newArr := []int{}
	stop := 0
	for i := 0; i < len(arr); i++ { // +-
		errorCode := 0
		stop--
		if i+2 < len(arr) && (arr[i] >= 0 || arr[i] < -6) && (arr[i+1] == -1 || arr[i+1] == -2) && (arr[i+2] >= 0 || arr[i+2] < -6) {
			if len(newArr) > 0 && (newArr[len(newArr)-1] >= 0 || newArr[len(newArr)-1] < -6) && stop == 1 {
				newArr[len(newArr)-1], errorCode = calcOper(newArr[len(newArr)-1], arr[i+1], arr[i+2])
				i++
			} else {
				var temp int
				temp, errorCode = calcOper(arr[i], arr[i+1], arr[i+2])
				newArr = append(newArr, temp)
				i++
			}
			stop = 2
		} else if stop <= 0 {
			newArr = append(newArr, arr[i])
		}
		if errorCode != 0 {
			return []int{}, errorCode
		}
	}
	return newArr, 0
}

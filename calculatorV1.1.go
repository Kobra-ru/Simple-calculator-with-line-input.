package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	/*
		Лист возможных дополнений/улучшений/уточнения:
			1)Ошибка огромных чисел даже не предупреждает
			2)Малое кол-во операторов
			3)Ошибочная ошибка при вводе оператора **
			4)Константы (число пи и тд.)
			5)Обработка букв насколько нужна большой вопрос?
			6)Оптимизация?Пару if/циклов удалить ну хз

	*/
	en := true
	line := ""
	fmt.Println("Здравствуйте, вы попали в калькулятор, чтобы выйти пропишите exit")
	fmt.Println("Числа: целые или дробные (разделители . или ,)\nДоступные операторы: + - * / ( )")
	for en {
		line = ""
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line = scanner.Text()
		if line == "exit" {
			en = false
		} else {
			var sumF float64
			var sumI int64
			var isInt bool
			var errorCode int
			sumF, sumI, isInt, errorCode = calc(line)
			//fmt.Printf("Дебаг сообщения %f %d %d %t\n", sumF, sumI, errorCode, isInt) //test
			if errorCode == 0 {
				if isInt {
					fmt.Printf("Ответ: %d\n", sumI)
				} else {
					if sumF-math.Floor(sumF) == 0.000000 {
						fmt.Printf("Ответ: %d\n", int(sumF))
					} else {
						count := 1
						for i := 10; i < 1000000; i = i * 10 {
							if sumF-math.Floor(sumF*float64(i))/float64(i) == 0.000000 {
								break
							}
							count++
						}
						switch count {
						case 1:
							sumF = math.Round(sumF*10) / 10
							fmt.Printf("Ответ: %0.1f\n", sumF)
						case 2:
							sumF = math.Round(sumF*100) / 100
							fmt.Printf("Ответ: %0.2f\n", sumF)
						case 3:
							sumF = math.Round(sumF*1000) / 1000
							fmt.Printf("Ответ: %0.3f\n", sumF)
						case 4:
							sumF = math.Round(sumF*10000) / 10000
							fmt.Printf("Ответ: %0.4f\n", sumF)
						case 5:
							sumF = math.Round(sumF*100000) / 100000
							fmt.Printf("Ответ: %0.5f\n", sumF)
						case 6:
							sumF = math.Round(sumF*1000000) / 1000000
							fmt.Printf("Ответ: %0.6f\n", sumF)
						}
					}
				}
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
				default:
					fmt.Println("Error: Неизвестная ошибка")
				}
				//5 код зарезервирован для системного использования
			}

		}
	}
}
func stringToActionInt(line string) ([]int64, int) { // обрабатывает ошибки Cинтаксическая, Неподдерживаемый оператор или Неподдерживаемый символ, Нет входных данных
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )

	var arr []int64 // основной массив
	var temp string
	var oldSym rune
	if line != "" {
		for i := 0; i < len(line); i++ {
			sym := rune(line[i])
			if sym == '+' || sym == '-' || sym == '*' || sym == '/' || sym == '(' || sym == ')' {
				if temp != "" {
					var tempN int64
					errorCode := 0
					tempN, errorCode = stringToInt(temp)
					arr = append(arr, tempN)
					temp = ""
					if errorCode != 0 {
						return []int64{}, errorCode
					}
				}
				switch sym {
				case '+':
					arr = append(arr, -1)
				case '-':
					arr = append(arr, -2)
				case '*':
					arr = append(arr, -3)
				case '/':
					return []int64{}, 5
				case '(':
					arr = append(arr, -5)
				case ')':
					arr = append(arr, -6)
				}

			} else if oldSym == ' ' && isNumber(rune(sym)) == true && temp != "" {
				return arr, 1
			} else if isNumber(rune(sym)) == true {
				temp = temp + string(sym)
			} else if sym == '.' || sym == ',' {
				return []int64{}, 5
			} else if sym != ' ' {
				return arr, 2
			}
			oldSym = rune(line[i])
		}
		if temp != "" {
			var tempN int64
			errorCode := 0
			tempN, errorCode = stringToInt(temp)
			arr = append(arr, tempN)
			temp = ""
			if errorCode != 0 {
				return []int64{}, errorCode
			}
		}
	}
	if len(arr) == 0 {
		return []int64{}, 3
	}
	return arr, 0
}
func stringToActionF(line string) ([]float64, int) { // обрабатывает ошибки Cинтаксическая, Неподдерживаемый оператор или Неподдерживаемый символ, Нет входных данных
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )

	var arr []float64 // основной массив
	var temp string
	var oldSym rune
	count := 0
	if line != "" {
		for i := 0; i < len(line); i++ {
			sym := rune(line[i])
			if sym == '+' || sym == '-' || sym == '*' || sym == '/' || sym == '(' || sym == ')' {
				if temp != "" {
					errorCode := 0

					switch count {
					case 0:
						var tempI int64
						tempI, errorCode = stringToInt(temp)
						arr = append(arr, float64(tempI))
					case 1:
						var tempF float64
						tempF, errorCode = stringToF(temp)
						arr = append(arr, tempF)
					default:
						return []float64{}, 2
					}
					temp = ""
					count = 0
					if errorCode != 0 {
						return []float64{}, errorCode
					}
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
			} else if sym == ',' || sym == '.' {
				temp = temp + "."
				count++
			} else if sym != ' ' {
				return arr, 2
			}
			oldSym = rune(line[i])
		}
		if temp != "" {
			errorCode := 0
			switch count {
			case 0:
				var tempI int64
				tempI, errorCode = stringToInt(temp)
				arr = append(arr, float64(tempI))
			case 1:
				var tempF float64
				tempF, errorCode = stringToF(temp)
				arr = append(arr, tempF)
			default:
				return []float64{}, 2
			}
			temp = ""
			count = 0
			if errorCode != 0 {
				return []float64{}, errorCode
			}
		}
	}
	if len(arr) == 0 {
		return []float64{}, 3
	}
	return arr, 0
}
func calc(line string) (float64, int64, bool, int) { // обрабатывает ошибки Cинтаксическая, Деление на ноль
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )
	var arrInt []int64
	var arrFloat []float64
	isInt := true
	errorCode := 0
	arrInt, errorCode = stringToActionInt(line)
	if errorCode == 5 {
		errorCode = 0
		arrFloat, errorCode = stringToActionF(line)
		isInt = false
	}
	if errorCode != 0 {
		return 0, 0, true, errorCode
	}

	var sumI int64
	var sumF float64
	sumI = 0
	sumF = 0
	if isInt {
		sumI, errorCode = actionCalcInt(arrInt)
	} else {
		sumF, errorCode = actionCalcFloat(arrFloat)
	}
	if errorCode != 0 {
		return 0, 0, true, errorCode
	}

	if isInt {
		if sumI <= -7 { // 				декодирования минусовых ответов int
			sumI = sumI + 7
		}
	} else {
		if sumF <= -7 { // 				декодирования минусовых ответов float
			sumF = sumF + 7
		}
	}

	return sumF, sumI, isInt, errorCode
}
func actionCalcFloat(arr []float64) (float64, int) {
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )
	//fmt.Println("Тест массив:", arr) //test
	errorCode := 0

	arr, errorCode = recursionOfStaplesF(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	arr = unaryMinusF(arr)

	arr, errorCode = multiplicationOrDivisionF(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	arr, errorCode = plusOrMinusF(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	if len(arr) == 0 || (arr[0] >= -6 && arr[0] < 0) || len(arr) > 1 { //ловля ошибок
		return 0, 1
	}

	return arr[0], 0
}
func actionCalcInt(arr []int64) (int64, int) {
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )
	//fmt.Println("Тест массив:", arr) //test
	errorCode := 0

	arr, errorCode = recursionOfStaplesInt(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	arr = unaryMinusInt(arr)

	arr, errorCode = multiplicationOrDivisionInt(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	arr, errorCode = plusOrMinusInt(arr)
	if errorCode != 0 {
		return 0, errorCode
	}

	if len(arr) == 0 || (arr[0] >= -6 && arr[0] < 0) || len(arr) > 1 { //ловля ошибок
		return 0, 1
	}

	return arr[0], 0
}
func calcOperInt(number1 int64, oper int64, number2 int64) (int64, int) {
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )
	if number1 <= -7 { //нормалезуем
		number1 = number1 + 7
	}
	if number2 <= -7 {
		number2 = number2 + 7
	}
	var sum int64
	sum = 0
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
		sum -= 7
	}
	return sum, 0
}
func calcOperF(number1 float64, oper int, number2 float64) (float64, int) {
	// -1 = +, -2 = -, -3 = *, -4 = /, -5 = (, -6 = )
	if number1 <= -7 { //нормалезуем
		number1 = number1 + 7
	}
	if number2 <= -7 {
		number2 = number2 + 7
	}
	var sum float64
	sum = 0
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
	sum = math.Round(sum*100000000) / 100000000
	if sum < 0 { //кодируем
		sum -= 7
	}
	return sum, 0
}
func isNumber(str rune) bool {
	if str == '0' || str == '1' || str == '2' || str == '3' || str == '4' || str == '5' || str == '6' || str == '7' || str == '8' || str == '9' {
		return true
	}
	return false
}
func stringToInt(str string) (int64, int) {
	var result int64
	result = 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '0':
		case '1':
			result = result + 1*int64(math.Pow(10, float64(len(str)-1-i)))
		case '2':
			result = result + 2*int64(math.Pow(10, float64(len(str)-1-i)))
		case '3':
			result = result + 3*int64(math.Pow(10, float64(len(str)-1-i)))
		case '4':
			result = result + 4*int64(math.Pow(10, float64(len(str)-1-i)))
		case '5':
			result = result + 5*int64(math.Pow(10, float64(len(str)-1-i)))
		case '6':
			result = result + 6*int64(math.Pow(10, float64(len(str)-1-i)))
		case '7':
			result = result + 7*int64(math.Pow(10, float64(len(str)-1-i)))
		case '8':
			result = result + 8*int64(math.Pow(10, float64(len(str)-1-i)))
		case '9':
			result = result + 9*int64(math.Pow(10, float64(len(str)-1-i)))
		default:
			return 0, 2
		}
	}
	return result, 0
}
func stringToF(str string) (float64, int) {
	var result float64
	result, errorCode := strconv.ParseFloat(str, 64)
	if errorCode != nil {
		return 0, 2
	}

	return result, 0
}
func recursionOfStaplesInt(arr []int64) ([]int64, int) {
	var newArr []int64
	var tempArr []int64
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
			var temp int64
			temp, errorCode = actionCalcInt(tempArr)
			newArr = append(newArr, temp)
			tempArr = []int64{}
		}
		if flag == true && quantity != 0 && !(quantity == 1 && arr[i] == -5) {
			tempArr = append(tempArr, arr[i])
		}
		if errorCode != 0 {
			return []int64{}, errorCode
		}
	}
	if len(newArr) == 0 || quantity != 0 { //ловля ошибок
		return []int64{}, 1
	}
	return newArr, 0
}
func recursionOfStaplesF(arr []float64) ([]float64, int) {
	var newArr []float64
	var tempArr []float64
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
			var temp float64
			temp, errorCode = actionCalcFloat(tempArr)
			newArr = append(newArr, temp)
			tempArr = []float64{}
		}
		if flag == true && quantity != 0 && !(quantity == 1 && arr[i] == -5) {
			tempArr = append(tempArr, arr[i])
		}
		if errorCode != 0 {
			return []float64{}, errorCode
		}
	}
	if len(newArr) == 0 || quantity != 0 { //ловля ошибок
		return []float64{}, 1
	}
	return newArr, 0
}
func unaryMinusInt(arr []int64) []int64 {
	newArr := []int64{}
	for i := 0; i < len(arr); i++ { //					унарный -

		if i == 0 && i+1 < len(arr) && arr[i] == -2 && (arr[i+1] >= 0 || arr[i+1] <= -7) {
			if arr[i+1] == 0 {
				newArr = append(newArr, 0)
			} else {
				newArr = append(newArr, arr[i+1]*-1-7)
			}
			i++
		} else if i > 0 && i+1 < len(arr) && arr[i] == -2 && arr[i-1] <= -1 && arr[i-1] >= -5 && (arr[i+1] >= 0 || arr[i+1] <= -7) {
			if arr[i+1] == 0 {
				newArr = append(newArr, 0)
			} else {
				newArr = append(newArr, arr[i+1]*-1-7)
			}
			i++
		} else {
			newArr = append(newArr, arr[i])
		}

	}
	return newArr
}
func unaryMinusF(arr []float64) []float64 {
	newArr := []float64{}
	for i := 0; i < len(arr); i++ { //					унарный -

		if i == 0 && i+1 < len(arr) && arr[i] == -2 && (arr[i+1] >= 0 || arr[i+1] <= -7) {
			if arr[i+1] == 0 {
				newArr = append(newArr, 0)
			} else {
				newArr = append(newArr, arr[i+1]*-1-7)
			}
			i++
		} else if i > 0 && i+1 < len(arr) && arr[i] == -2 && arr[i-1] <= -1 && arr[i-1] >= -5 && (arr[i+1] >= 0 || arr[i+1] <= -7) {
			if arr[i+1] == 0 {
				newArr = append(newArr, 0)
			} else {
				newArr = append(newArr, arr[i+1]*-1-7)
			}
			i++
		} else {
			newArr = append(newArr, arr[i])
		}

	}
	return newArr
}
func multiplicationOrDivisionInt(arr []int64) ([]int64, int) {
	newArr := []int64{}
	stop := 0
	for i := 0; i < len(arr); i++ { // умножения деления
		errorCode := 0
		stop--
		if i+2 < len(arr) && (arr[i] >= 0 || arr[i] <= -7) && (arr[i+1] == -3 || arr[i+1] == -4) && (arr[i+2] >= 0 || arr[i+2] <= -7) {
			if len(newArr) > 0 && (newArr[len(newArr)-1] >= 0 || newArr[len(newArr)-1] <= -7) && stop == 1 {
				newArr[len(newArr)-1], errorCode = calcOperInt(newArr[len(newArr)-1], arr[i+1], arr[i+2])
				i++
			} else {
				var temp int64
				temp, errorCode = calcOperInt(arr[i], arr[i+1], arr[i+2])
				newArr = append(newArr, temp)
				i++
			}
			stop = 2
		} else if stop <= 0 {
			newArr = append(newArr, arr[i])
		}
		if errorCode != 0 {
			return []int64{}, errorCode
		}

	}
	return newArr, 0
}
func multiplicationOrDivisionF(arr []float64) ([]float64, int) {
	newArr := []float64{}
	stop := 0
	for i := 0; i < len(arr); i++ { // умножения деления
		errorCode := 0
		stop--
		if i+2 < len(arr) && (arr[i] >= 0 || arr[i] <= -7) && (arr[i+1] == -3 || arr[i+1] == -4) && (arr[i+2] >= 0 || arr[i+2] <= -7) {
			if len(newArr) > 0 && (newArr[len(newArr)-1] >= 0 || newArr[len(newArr)-1] <= -7) && stop == 1 {
				newArr[len(newArr)-1], errorCode = calcOperF(newArr[len(newArr)-1], int(arr[i+1]), arr[i+2])
				i++
			} else {
				var temp float64
				temp, errorCode = calcOperF(arr[i], int(arr[i+1]), arr[i+2])
				newArr = append(newArr, temp)
				i++
			}
			stop = 2
		} else if stop <= 0 {
			newArr = append(newArr, arr[i])
		}
		if errorCode != 0 {
			return []float64{}, errorCode
		}

	}
	return newArr, 0
}
func plusOrMinusInt(arr []int64) ([]int64, int) {
	newArr := []int64{}
	stop := 0
	for i := 0; i < len(arr); i++ { // +-
		errorCode := 0
		stop--
		if i+2 < len(arr) && (arr[i] >= 0 || arr[i] <= -7) && (arr[i+1] == -1 || arr[i+1] == -2) && (arr[i+2] >= 0 || arr[i+2] <= -7) {
			if len(newArr) > 0 && (newArr[len(newArr)-1] >= 0 || newArr[len(newArr)-1] <= -7) && stop == 1 {
				newArr[len(newArr)-1], errorCode = calcOperInt(newArr[len(newArr)-1], arr[i+1], arr[i+2])
				i++
			} else {
				var temp int64
				temp, errorCode = calcOperInt(arr[i], arr[i+1], arr[i+2])
				newArr = append(newArr, temp)
				i++
			}
			stop = 2
		} else if stop <= 0 {
			newArr = append(newArr, arr[i])
		}
		if errorCode != 0 {
			return []int64{}, errorCode
		}
	}
	return newArr, 0
}
func plusOrMinusF(arr []float64) ([]float64, int) {
	newArr := []float64{}
	stop := 0
	for i := 0; i < len(arr); i++ { // +-
		errorCode := 0
		stop--
		if i+2 < len(arr) && (arr[i] >= 0 || arr[i] <= -7) && (arr[i+1] == -1 || arr[i+1] == -2) && (arr[i+2] >= 0 || arr[i+2] <= -7) {
			if len(newArr) > 0 && (newArr[len(newArr)-1] >= 0 || newArr[len(newArr)-1] <= -7) && stop == 1 {
				newArr[len(newArr)-1], errorCode = calcOperF(newArr[len(newArr)-1], int(arr[i+1]), arr[i+2])
				i++
			} else {
				var temp float64
				temp, errorCode = calcOperF(arr[i], int(arr[i+1]), arr[i+2])
				newArr = append(newArr, temp)
				i++
			}
			stop = 2
		} else if stop <= 0 {
			newArr = append(newArr, arr[i])
		}
		if errorCode != 0 {
			return []float64{}, errorCode
		}
	}
	return newArr, 0
}

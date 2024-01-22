package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Недопустимая операция")
	}
}

func main() {
	for {

		var input string
		fmt.Print("Введите выражение: ")

		myscanner := bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		input = myscanner.Text()

		// Разделяем введенную строку на операнды и оператор
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			panic("Недопустимое выражение превышено количество значений или мало значений")
		}

		aStr, operator, bStr := parts[0], parts[1], parts[2]

		// Проверяем, являются ли операнды только римскими числами или арабскими цифрами
		if (isArabicNumber(aStr) && isArabicNumber(bStr)) || (isRomanNumber(aStr) && isRomanNumber(bStr)) {
		} else {
			panic(" операнды не только римские числа или арабские числа")
		}
		// Преобразуем операнды в арабские числа если они романские
		var a, b int
		if isRomanNumber(aStr) {
			a = convertRomanToArabic(aStr)
			b = convertRomanToArabic(bStr)
		} else {
			a, _ = strconv.Atoi(aStr)
			b, _ = strconv.Atoi(bStr)
		}
		// Проверяем попадает ли в диапазон чисел от 1 до 10
		if (a <= 10) && (a > 0) && (b <= 10) && (b > 0) {
		} else {
			panic("числа не попадают  в диапазон чисел от 1 до 10 ")
		}

		// Выполняем операцию и выводим результат
		result := calculate(a, b, operator)
		if isRomanNumber(aStr) {
			fmt.Println(convertArabicToRoman(result))
		} else {
			fmt.Println(result)
		}
	}
}
func convertRomanToArabic(roman string) int {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return romanNumerals[roman]
}
func convertArabicToRoman(arabic int) string {
	romanNumerals := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16: "XVI",
		17: "XVII",
		18: "XVIII",
		19: "XIX",
		20: "XX",
	}
	v := romanNumerals[arabic]
	if v == "" {
		panic("числа недопустимо для римских цифр")
	}
	return v
}

func isArabicNumber(str string) bool {
	// Проверяем, содержит ли строка только символы, соответствующие арабским цифрам
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func isRomanNumber(str string) bool {
	// Проверяем, содержит ли строка только символы, соответствующие римским цифрам
	romanNumerals := []string{"I", "V", "X", "L", "C", "D", "M"}

	for _, char := range str {
		if !strings.ContainsRune(strings.Join(romanNumerals, ""), char) {
			return false
		}
	}
	return true
}

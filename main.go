package main

import (
	"fmt"
	"strconv"
)

func main() {

	var firstNum, secondNum, symbol string
	var exit int
	cheker1 := "Я не придумал как это сделать по-другому, поэтому я сделал эту переменную чтобы проверять не введено-ли что-то еще после примера. Обязательно придумаю позже"
	cheker2 := "Я не придумал как это сделать по-другому, поэтому я сделал эту переменную чтобы проверять не введено-ли что-то еще после примера. Обязательно придумаю позже"
	fmt.Scanln(&firstNum, &symbol, &secondNum, &cheker1)

	switch {
	case cheker1 != cheker2:

		fmt.Println("Ошибка: Введено больше данных(Программа принимает только пример вида '1 + 1')")

	case whatIs(symbol) == "symbol" && whatIs(firstNum) == "arabic" && whatIs(secondNum) == "arabic":

		fmt.Println(counting(firstNum, symbol, secondNum))

	case whatIs(symbol) == "symbol" && whatIs(firstNum) == "rome" && whatIs(secondNum) == "rome":

		fmt.Println(arabicToRome(counting(romeToArabic(firstNum), symbol, romeToArabic(secondNum))))

	case (whatIs(symbol) != "symbol") || (whatIs(firstNum) != "rome" && whatIs(firstNum) != "arabic") || (whatIs(secondNum) != "rome" && whatIs(secondNum) != "arabic"):

		fmt.Println("Ошибка: На вход поданы несоответствующие числа(0 < a <= 10), не числа, не целые числа или использован неправильный оператор")

	case (whatIs(firstNum) == "rome" && whatIs(secondNum) == "arabic") || whatIs(firstNum) == "arabic" && whatIs(secondNum) == "rome":

		fmt.Println("Ошибка: Одновременно используются разные системы счисления.")
	}

	fmt.Println("Если хотите продолжить - введите 1. При любом другом вводе программа будет завершена")
	fmt.Scan(&exit)
	if exit == 1 {
		main()
	}

}

func whatIs(a string) string {
	arabicNum := [11]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	romeNum := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	chekSymbol := [4]string{"+", "-", "*", "/"}

	var cheker string

	for _, v := range arabicNum {
		switch {
		case v == a:
			cheker = "arabic"
		}
	}

	for _, v := range romeNum {
		switch {
		case v == a:
			cheker = "rome"
		}
	}

	for _, v := range chekSymbol {
		switch {
		case v == a:
			cheker = "symbol"
		}
	}

	return cheker
}

func counting(firstNum, symbol, secondNum string) int {

	trueFirstNum, _ := strconv.Atoi(firstNum)
	trueSecondNum, _ := strconv.Atoi(secondNum)

	var result int

	switch {
	case symbol == "+":
		result = trueFirstNum + trueSecondNum
	case symbol == "-":
		result = trueFirstNum - trueSecondNum
	case symbol == "*":
		result = trueFirstNum * trueSecondNum
	case symbol == "/":
		result = trueFirstNum / trueSecondNum
	}

	return result
}
func romeToArabic(romeNum string) string {

	var result string

	switch romeNum {
	case "I":
		result = "1"
	case "II":
		result = "2"
	case "III":
		result = "3"
	case "IV":
		result = "4"
	case "V":
		result = "5"
	case "VI":
		result = "6"
	case "VII":
		result = "7"
	case "VIII":
		result = "8"
	case "IX":
		result = "9"
	case "X":
		result = "10"
	}

	return result
}

func arabicToRome(arabicNum int) string {

	var tensArabic, onesArabic int
	var tensRome, onesRome string
	var result string

	tensArabic = arabicNum / 10
	onesArabic = arabicNum % 10

	switch tensArabic {
	case 1:
		tensRome = "X"
	case 2:
		tensRome = "XX"
	case 3:
		tensRome = "XXX"
	case 4:
		tensRome = "XL"
	case 5:
		tensRome = "L"
	case 6:
		tensRome = "LX"
	case 7:
		tensRome = "LXX"
	case 8:
		tensRome = "LXXX"
	case 9:
		tensRome = "XC"
	case 10:
		tensRome = "C"
	}

	switch onesArabic {
	case 1:
		onesRome = "I"
	case 2:
		onesRome = "II"
	case 3:
		onesRome = "III"
	case 4:
		onesRome = "IV"
	case 5:
		onesRome = "V"
	case 6:
		onesRome = "VI"
	case 7:
		onesRome = "VII"
	case 8:
		onesRome = "VIII"
	case 9:
		onesRome = "IX"
	case 0:
		onesRome = ""
	}

	result = tensRome + onesRome

	if arabicNum < 1 {
		result = "Ошибка: В римской системе счисления нет отрицательных чисел и нуля"
	}

	return result
}

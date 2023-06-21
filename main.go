package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		cheker1   string = "Я не придумал как это сделать по-другому, поэтому я сделал эту переменную чтобы проверять не введено-ли что-то еще после примера. Обязательно придумаю позже"
		cheker2   string = "Я не придумал как это сделать по-другому, поэтому я сделал эту переменную чтобы проверять не введено-ли что-то еще после примера. Обязательно придумаю позже"
		error1    string = "Ошибка: Введено больше данных(Программа принимает только пример вида '1 + 1')"
		error2    string = "Ошибка: На вход поданы несоответствующие числа [1:10], не числа, не целые числа или использован неправильный оператор"
		error3    string = "Ошибка: Одновременно используются разные системы счисления."
		error4    string = "Ошибка: В римской системе счисления нет отрицательных чисел и нуля"
		firstNum  string
		secondNum string
		symbol    string
	)

	fmt.Scanln(&firstNum, &symbol, &secondNum, &cheker1)

	var answer int = counting(firstNum, symbol, secondNum)

	switch {
	case cheker1 != cheker2:

		fmt.Println(error1)

	case (whatIs(symbol) != "symbol") || (whatIs(firstNum) != "rome" && whatIs(firstNum) != "arabic") || (whatIs(secondNum) != "rome" && whatIs(secondNum) != "arabic"):

		fmt.Println(error2)

	case (whatIs(firstNum) == "rome" && whatIs(secondNum) == "arabic") || whatIs(firstNum) == "arabic" && whatIs(secondNum) == "rome":

		fmt.Println(error3)

	case whatIs(symbol) == "symbol" && whatIs(firstNum) == "rome" && whatIs(secondNum) == "rome" && answer < 1:

		fmt.Println(error4)

	case whatIs(symbol) == "symbol" && whatIs(firstNum) == "arabic" && whatIs(secondNum) == "arabic":

		fmt.Println(answer)

	case whatIs(symbol) == "symbol" && whatIs(firstNum) == "rome" && whatIs(secondNum) == "rome" && answer > 0:

		fmt.Println(arabicToRome(answer))
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

	if whatIs(firstNum) == "rome" && whatIs(secondNum) == "rome" {
		firstNum = romeToArabic(firstNum)
		secondNum = romeToArabic(secondNum)
	}

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

	return result
}

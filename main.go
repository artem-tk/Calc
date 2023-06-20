package main

import (
	"fmt"
	"strconv"
)

func main() {

	var firstNum, secondNum, symbol string

	fmt.Scanln(&firstNum, &symbol, &secondNum)

	switch {
	case whatIs(symbol) == "symbol" && whatIs(firstNum) == "arabic" && whatIs(secondNum) == "arabic":
		fmt.Println("Оба числа арабские, символ правильный. Произвожу вычисление...")
		fmt.Println(counting(firstNum, symbol, secondNum))
	case whatIs(symbol) == "symbol" && whatIs(firstNum) == "rome" && whatIs(secondNum) == "rome":
		fmt.Println("Оба числа римские, символ правильный. Произвожу вычисление...")
		fmt.Println(arabicToRome(counting(romeToArabic(firstNum), symbol, romeToArabic(secondNum))))
	}

}

func whatIs(a string) string {
	arabicNum := [11]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
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
	var result string

	switch arabicNum {
	case 1:
		result = "I"
	case 2:
		result = "II"
	case 3:
		result = "III"
	case 4:
		result = "IV"
	case 5:
		result = "V"
	case 6:
		result = "VI"
	case 7:
		result = "VII"
	case 8:
		result = "VIII"
	case 9:
		result = "IX"
	case 10:
		result = "X"
	}

	return result
}

package main

import (
	rome_digit "awesomeProject/operat"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение")
	text, _ := reader.ReadString('\n')
	replaceAll := strings.ReplaceAll(text, " ", "")
	trim := strings.Trim(replaceAll, "\n")
	array := RegSplit(trim, "[-+/*]")
	if len(array) > 2 {
		fmt.Println("Должно быть только 2 числа")
		os.Exit(1)

	}
	x := array[0]
	y := array[1]
	operationValue := getOperationValue(text)
	if isDigit(x) && isDigit(y) {
		x1, _ := strconv.Atoi(x)
		y1, _ := strconv.Atoi(y)
		if (x1 <= 0 || x1 > 10) || (y1 <= 0 || y1 > 10) {
			fmt.Println("Числа должны быть в диапазоне 1 - 10")
			os.Exit(1)
		}
		result := rome_digit.GetResult(x1, operationValue, y1)
		fmt.Println(result)

	} else if isRome(x) && isRome(y) {

		x1 := romeToInt(x)
		y1 := romeToInt(y)
		if (x1 <= 0 || x1 > 10) || (y1 <= 0 || y1 > 10) {
			fmt.Println("Числа должны быть в диапазоне 1 - 10")
			os.Exit(1)
		}
		result := rome_digit.GetResult(x1, operationValue, y1)
		if result <= 0 {
			fmt.Println("Ошибка, в римских числах нет 0 и отрицательных значений")
			os.Exit(1)
		}
		res := intToRome(result)
		fmt.Println(res)

	} else {
		fmt.Println("Неверный формат")
	}
}

func getOperationValue(text string) string {
	var val string
	if strings.Contains(text, "*") {
		val = "*"
	} else if strings.Contains(text, "+") {
		val = "+"
	} else if strings.Contains(text, "-") {
		val = "-"
	} else if strings.Contains(text, "/") {
		val = "/"
	} else {
		fmt.Println("Данная операция не поддерживается")
		os.Exit(1)
	}
	return val
}
func isDigit(s string) bool {
	var flag bool
	atoi, _ := strconv.Atoi(s)
	if atoi >= 1 {
		flag = true
	} else {

		flag = false
	}
	return flag
}

func isRome(s string) bool {
	var flag bool
	for key := range num {
		if key == s {
			flag = true
			break
		}
	}
	return flag
}
func romeToInt(s string) int {
	var romeToInt int
	for key, value := range num {
		if key == s {
			romeToInt = value
			return romeToInt
		}
	}
	return romeToInt
}

func intToRome(i int) string {
	var res string
	switch i {
	case 1:
		res = "I"
	case 2:
		res = "II"
	case 3:
		res = "III"
	case 4:
		res = "IV"
	case 5:
		res = "V"
	case 6:
		res = "VI"
	case 7:
		res = "VII"
	case 8:
		res = "VIII"
	case 9:
		res = "IX"
	case 10:
		res = "X"
	case 11:
		res = "XI"
	case 12:
		res = "XII"
	case 13:
		res = "XIII"
	case 14:
		res = "XIV"
	case 15:
		res = "XV"
	case 16:
		res = "XVI"
	case 17:
		res = "XVIi"
	case 18:
		res = "XVIII"
	case 19:
		res = "XIX"
	case 20:
		res = "XX"
	case 21:
		res = "XXI"
	case 22:
		res = "XXII"
	case 23:
		res = "XXIII"
	case 24:
		res = "XXIV"
	case 25:
		res = "XXV"
	case 26:
		res = "XXVI"
	case 27:
		res = "XXVII"
	case 28:
		res = "XXVIII"
	case 29:
		res = "XXIX"
	case 30:
		res = "XXX"
	case 31:
		res = "XXXI"
	case 32:
		res = "XXXII"
	case 33:
		res = "XXXIII"
	case 34:
		res = "XXXIV"
	case 35:
		res = "XXXIV"
	case 36:
		res = "XXXVI"
	case 37:
		res = "XXXVII"
	case 38:
		res = "XXXVIII"
	case 39:
		res = "XXXIX"
	case 40:
		res = "XL"
	case 41:
		res = "XLI"
	case 42:
		res = "XLII"
	case 43:
		res = "XLIII"
	case 44:
		res = "XLIV"
	case 45:
		res = "XLV"
	case 46:
		res = "XLVI"
	case 47:
		res = "XLVII"
	case 48:
		res = "XLVIII"
	case 49:
		res = "XLIX"
	case 50:
		res = "L"
	case 51:
		res = "LI"
	case 52:
		res = "LII"
	case 53:
		res = "LIII"
	case 54:
		res = "LIV"
	case 55:
		res = "LV"
	case 56:
		res = "LVI"
	case 57:
		res = "LVII"
	case 58:
		res = "LVIII"
	case 59:
		res = "LIX"
	case 60:
		res = "LX"
	case 61:
		res = "LXI"
	case 62:
		res = "LXII"
	case 63:
		res = "LXIII"
	case 64:
		res = "LXIV"
	case 65:
		res = "LXV"
	case 66:
		res = "LXVI"
	case 67:
		res = "LXVIII"
	case 68:
		res = "LXIX"
	case 69:
		res = "LXX"
	case 70:
		res = "LXX"
	case 71:
		res = "LXXI"
	case 72:
		res = "LXXII"
	case 73:
		res = "LXXIII"
	case 74:
		res = "LXXIV"
	case 75:
		res = "LXXV"
	case 76:
		res = "LXXVI"
	case 77:
		res = "LXXVII"
	case 78:
		res = "LXXVIII"
	case 79:
		res = "LXXIX"
	case 80:
		res = "LXXX"
	case 81:
		res = "LXXXI"
	case 82:
		res = "LXXXII"
	case 83:
		res = "LXXXIII"
	case 84:
		res = "LXXXIV"
	case 85:
		res = "LXXXV"
	case 86:
		res = "LXXXVI"
	case 87:
		res = "LXXXVII"
	case 88:
		res = "LXXXVIII"
	case 89:
		res = "LXXXIX"
	case 90:
		res = "XC"
	case 91:
		res = "XCI"
	case 92:
		res = "XCII"
	case 93:
		res = "XCIII"
	case 94:
		res = "XCIV"
	case 95:
		res = "XCV"
	case 96:
		res = "XCVI"
	case 97:
		res = "XCVII"
	case 98:
		res = "XCVIII"
	case 99:
		res = "XCIX"
	case 100:
		res = "C"
	default:
		fmt.Println("Результат превзошел ожидания")
	}
	return res
}

var num = map[string]int{
	"I":     1,
	"II":    2,
	"III":   3,
	"IV":    4,
	"V":     5,
	"VI":    6,
	"VII":   7,
	"VIII":  8,
	"IX":    9,
	"X":     10,
	"XI":    11,
	"XII":   12,
	"XIII":  13,
	"XIV":   14,
	"XV":    15,
	"XVI":   16,
	"XVII":  17,
	"XVIII": 18,
	"XIX":   19,
	"XX":    20,
}

func RegSplit(text string, delimeter string) []string {
	var a []string
	if strings.Contains(text, "*") || strings.Contains(text, "/") || strings.Contains(text, "+") || strings.Contains(text, "-") {
		reg := regexp.MustCompile(delimeter)
		text1 := strings.TrimSpace(text)
		indexes := reg.FindAllStringIndex(text1, -1)
		laststart := 0
		a := make([]string, len(indexes)+1)
		for i, element := range indexes {
			a[i] = text[laststart:element[0]]
			laststart = element[1]
		}
		a[len(indexes)] = text[laststart:len(text)]
		return a
	}
	fmt.Println("Неверный знак")

	return a
}

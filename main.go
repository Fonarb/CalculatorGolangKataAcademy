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
			fmt.Println("Ошибка, в римских числах нет 0")
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
	if atoi > 1 {
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
	var str string
	for key, value := range num {
		if value == i {
			str = key
			return str
		}
	}
	return str
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

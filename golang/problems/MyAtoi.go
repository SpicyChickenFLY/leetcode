package main

import "fmt"
import "strconv"

func myAtoi(str string) int {
	var status = [][] int {
		{0, 3, 1, 4, 2, 4},
		{1, 1, 1, 1, 1, 1},
		{5, 5, 5, 2, 2, 5},
		{1, 1, 1, 4, 2, 1},
		{1, 1, 1, 4, 2, 1},
		{5, 5, 5, 5, 5, 5},
		{1, 1, 1, 4, 2, 1},
	}
	var currentStatus int = 0
	var resultString = "0"
	var sign bool = true
	for _, character := range str {

		switch {
			case character == [] rune (" ")[0] ://0
				currentStatus = status[currentStatus][0]
			case character == [] rune ("-")[0] : //1
				currentStatus = status[currentStatus][1]
			case character == [] rune ("0")[0]: //3
				currentStatus = status[currentStatus][3]
			case character >= [] rune ("1")[0] && character <= [] rune ("9")[0] : //4
				currentStatus = status[currentStatus][4]
			case character == [] rune ("+")[0]: //5
				currentStatus = status[currentStatus][5]
			default://2
				currentStatus = status[currentStatus][2]
		}

		var flag bool = true
		switch currentStatus {
		case 1:
			//fmt.Printf(" 1")
			return 0
		case 2:
			resultString += string(character)
		case 3:
			sign = false
		case 5:
			flag = false
		}
		//fmt.Printf("%2d->", currentStatus)
		if !flag {
			break
		}
	}
	result, _ := strconv.ParseInt(resultString, 10, 64)
	if sign && result > 2147483647 {
		return int(2147483647)
	}
	if !sign {
		if result > 2147483648 {
			return int(-2147483648)
		}
		result *= -1
	}
    return int(result)
}

func main(){
	var inputSet = [] string {
		"42",
		"0",
		"0001s",
		"    -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
		"-",
		"+",
		"+1",
	}
	for index, input := range inputSet {
		fmt.Printf("\n%2d: %d\n", index, myAtoi(input))
	}
}
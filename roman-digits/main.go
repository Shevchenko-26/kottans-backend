package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digitLength = 7
var digitWidth = 5

func main() {
	if len(os.Args) != 2 {
		log.Println("Need only one input number")
		log.Println(os.Args)
		os.Exit(1)
	}
	number := os.Args[1]
	if _, err := strconv.ParseUint(number, 10, 64); err != nil {
		log.Println("Input argument is integer non-convertible: ", err.Error())
		os.Exit(1)
	}

	digs := prepareDigits()

	sb := strings.Builder{}
	for i := 0; i < 7; i++ {
		for _, num := range number {
			sb.WriteString(digs[int(num-48)][i] + " ")
		}
		if i != digitLength-1 {

			sb.WriteByte('\n')
		}
	}
	asteriskBorder := strings.Repeat("*", len(number)*(digitWidth+1)-1)

	fmt.Printf("%s\n%s\n%s", asteriskBorder, sb.String(), asteriskBorder)
}

func prepareDigits() map[int][]string {
	artNumbers := strings.Split(digits, "\n\n")
	numbers := make(map[int][]string)
	numberHeight := 0
	for _, digit := range artNumbers {
		for _, digitString := range strings.Split(digit, "\n") {
			if digitString == "" {
				continue
			}
			digitString = fmt.Sprintf("%-5s", digitString)
			numbers[numberHeight/digitLength] = append(numbers[numberHeight/digitLength], digitString)
			numberHeight++
		}
	}
	return numbers
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("ERROR! That number is not even number..")
	}

	return num, nil
}

func inputNum() {

	fmt.Print("Please write a number : ")
	reader := bufio.NewReader(os.Stdin)
	number1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	number1 = strings.TrimSpace(number1)
	number2, err := strconv.Atoi(number1)
	result, err := evenNum(number2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your number is :", result)
	}
}

func main() {
	inputNum()
}

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var expressionArray []string

func main() {
	fmt.Println("Processing...")
	twoHundred()
}

func twoHundred() {

	var sum int
	var textExpression string
	var loop int

	for {
		a := arrayCombination()
		sum = 0
		textExpression = ""
		loop = 0
		for _, v := range a {
			if loop == 0 {
				sum = v
				textExpression += strconv.Itoa(v)
				loop++
				continue
			}

			time.Sleep(5 * time.Millisecond)

			rand.Seed(time.Now().UnixNano())
			sign := 1 + rand.Intn(2-1+1)

			if sign <= 1 {
				sum -= v
				textExpression += "-" + strconv.Itoa(v)
			} else {
				sum += v
				textExpression += "+" + strconv.Itoa(v)
			}
		}

		if sum == 200 {
			//fmt.Println("200")
			ok := findExpression(textExpression)
			if !ok {
				expressionArray = append(expressionArray, textExpression)
				fmt.Println(textExpression, "=", sum)
			}
		}
	}

}

func arrayCombination() []int {
	var result []int
	var step int
	var b bool = false
	var s string
	var appended bool
	for i := 9; i >= 0; i-- {
		appended = false
		s += strconv.Itoa(i)

		time.Sleep(7 * time.Millisecond)

		if !b {
			rand.Seed(time.Now().UnixNano())
			step = rand.Intn(3)

			b = true
		}

		if step <= 1 {
			b = false
			t, _ := strconv.Atoi(s)
			result = append(result, t)
			appended = true
			s = ""
		}

		step--

		if i == 0 && !appended {
			t, _ := strconv.Atoi(s)
			result = append(result, t)
		}

	}

	return result
}

func findExpression(expression string) bool {
	for _, v := range expressionArray {
		if v == expression {
			return true
		}
	}

	return false
}

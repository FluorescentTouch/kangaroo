package main

// Kangaroo is simple program that calculates possibility of meeting of two kangaroos with determined starting speeds
// and coordinates. Program requires input string in format `x1 v1 x2 v2` divided by spaces.
// The answer would be printed on screen: YES, they will met in same time in same place, or NO, they wont met.

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"errors"
	"strconv"
)

var (
	x1, x2, v1, v2 int

	errorInvalidValue = errors.New("invalid input value")
	)

 const (
 	xMax = 10000
 	xMin = -10000

 	vMax = 10000
 	vMin = -10000
 	)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter values")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input = strings.TrimSuffix(input, "\n")
	x1,v1,x2,v2,err := validateValues(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	theyWillMet := kenGo(x1,v1,x2,v2)

	if theyWillMet {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}


func kenGo(x1 ,v1 ,x2 ,v2 int) bool {
	if v1 == v2 {
		if x1 == x2 {
			return true
		}
		return false
	}

	if x1 == x2 {
		if v1 == v2 {
			return true
		}
		return false
	}

	/*
	x1 = x01 + v1*t
	x2 = x02 + v2*t

	x1 = x2 ->

	t = (x01 - x02)/(v2 -v1)

	if t > 0 and t - int -> return true
	else return false

	division by zero excluded by primary conditions
	*/

	if (x1 - x2) % (v2 - v1) == 0 &&
		(x1 - x2) / (v2 - v1) > 0 {
			return true
	}
	return false
}

func validateValues(v string) (x1 ,v1 ,x2 ,v2 int, err error) {
	if len(v) == 0 {
		err = errorInvalidValue
		return
	}

	values := strings.Split(v, " ")
	if len(values) < 4 {
		err = errorInvalidValue
		return
	}
	x1, err = strconv.Atoi(values[0])
	if err != nil {
		err = errorInvalidValue
		return
	}
	v1, err = strconv.Atoi(values[1])
	if err != nil {
		err = errorInvalidValue
		return
	}
	x2, err = strconv.Atoi(values[2])
	if err != nil {
		err = errorInvalidValue
		return
	}
	v2, err = strconv.Atoi(values[3])
	if err != nil {
		err = errorInvalidValue
		return
	}

	switch {
	case x1 < xMin || x1 > xMax:
		err = errorInvalidValue
	case v1 < vMin || v1 > vMax:
		err = errorInvalidValue
	case x2 < xMin || x2 > xMax:
		err = errorInvalidValue
	case v2 < vMin || v2 > vMax:
		err = errorInvalidValue
	default:
	}
	return
}
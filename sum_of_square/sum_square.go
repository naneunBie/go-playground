package main

import (
	"fmt"
)

func main() {

	// first input will be initial sum of total test case
	var n int
	fmt.Print("Please input total test case : ")
	fmt.Scan(&n)

	var result []int

	process(n, 0, result)
}

func process(n int, counter int, result []int) {

	// will checking if total test case already 0 will be stop
	if n <= 0 {
		return
	}

	if n == counter {
		return
	}

	var num int
	fmt.Scan(&num)

	// only will be append the result if
	sum := sum_of_square(num)
	if num > 0 {
		result = append(result, sum)
	}

	// print the results
	if len(result) >= n {
		print_result(0, result)
	}

	process(n, counter+1, result)
}

func print_result(i int, result []int) {

	if i >= len(result) {
		return
	}

	fmt.Println(result[i])

	print_result(i+1, result)

}

func sum_of_square(num int) int {
	// validate if input number negative, will direct return 0
	if num == 0 {
		fmt.Println()
		return 0
	}

	var value int
	fmt.Scan(&value)

	if value > 0 {
		return value*value + sum_of_square(num-1)
	}

	return sum_of_square(num - 1)
}

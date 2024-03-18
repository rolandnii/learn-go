package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int, int) int {
	return func(first, second int) int {
		return first + second
	}
}

func main() {
	f := fibonacci()
	first := 0
	second := 0
	sum := 0
	//sum 0  if i = 1 sum = 1 but the old or first is 0
	for i := 0; i < 13; i++ {

		fmt.Println(sum)

		if i < 2 {
			first = 0
			sum = 1
			first, second = second, sum

		} else {
			sum = f(first, second)
			first, second = second, sum
		}

	}
}

package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
func main() {
	fmt.Printf("%d\n", Sum([]int{1, 2, 3}))
}

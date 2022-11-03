package main

import "fmt"

func Count(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			// Blocks on the operation
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}

func main() {
	const beers = 10
	fmt.Println("No bottles of beer on the wall")

	for i := range Count(1, beers) {
		fmt.Println("Pass it around, put one up,", i, "bottles of beer on the wall")
		// Pass it around, put one up, 1 bottles of beer on the wall
		// Pass it around, put one up, 2 bottles of beer on the wall
		// ...
		// Pass it around, put one up, 99 bottles of beer on the wall
	}

	fmt.Println(beers, "bottles of beer on the wall")
}

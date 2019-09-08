package main

import "fmt"

type number []int

func main() {
	numbers := number{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers.check()

}

func (n number) check() {
	for _, number := range n {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}

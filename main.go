package main

import "fmt"

func main() {
	b := []int{0, 0}
	test(b)
	test2(&b)
	fmt.Println(b)
}

func test(b []int) {
	b[0] = 1
}

func test2(b *[]int) {
	(*b)[1] = 2
}

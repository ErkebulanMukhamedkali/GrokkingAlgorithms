package main

import (
	"algos/chapter1"
	"log"
)

func main() {
	target := 9
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(chapter1.DumbSearch(target, input))
	log.Println(chapter1.BinarySearch(target, input))
}

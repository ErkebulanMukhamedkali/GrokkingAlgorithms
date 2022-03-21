package main

import (
	"algos/chapter1"
	"algos/chapter2"
	"log"
)

func main() {
	target := 9
	input := []int{1, 3, 5, 7, 9}
	log.Println(chapter1.DumbSearch(target, input))
	log.Println(chapter1.BinarySearch(target, input))

	input = []int{5, 3, 6, 2, 10}
	log.Println(chapter2.SelectionSortV2(input))
	log.Println(chapter2.SelectionSort(input))
}

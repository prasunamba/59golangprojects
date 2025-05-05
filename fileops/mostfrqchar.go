package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Mostfreqchar() {
	/* intslice := []int{2, 1, 3, 4, 2, 1, 3, 4, 5, 4}
	sum := getsubscribe(intslice)
	fmt.Println("sum", sum) */
	f, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	fremap := make(map[rune]int)
	for _, val := range string(data) {
		if val != ' ' && val != '\n' {
			fremap[val]++
		}
	}
	var result rune
	maxcount := -1
	for key, value := range fremap {
		if value > maxcount {
			maxcount = value
			result = key
		}
	}
	fmt.Printf("result %c ", result)

}

/* func getsubscribe(intslice []int) int {
	mapofintslice := make(map[int]int)
	for _, value := range intslice {
		mapofintslice[value]++
	}
	mostoccuredint := -1
	result := -1
	fmt.Println("", mapofintslice)
	for key, value := range mapofintslice {
		if mostoccuredint < value {
			mostoccuredint = value
			result = key
			fmt.Println("-", result)
		}
	}
	return result
} */

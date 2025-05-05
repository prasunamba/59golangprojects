package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Mostfreqword() {

	f, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	fremap := make(map[string]int)
	sentence := strings.Fields(string(data))
	for _, val := range sentence {
		if val != " " && val != "\n" {
			cleaned := strings.Trim(val, ".,!?;:\"'()[]{}") // remove punctuation
			fremap[cleaned]++
		}
	}
	var result []string

	maxcount := -1
	fmt.Println("", fremap)
	for _, value := range fremap {
		if value > maxcount {
			maxcount = value
		}
	}
	for key, value := range fremap {
		if value == maxcount {
			result = append(result, key)
		}
	}
	fmt.Println("result  ", result, maxcount)

}

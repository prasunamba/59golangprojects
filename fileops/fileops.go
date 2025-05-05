package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Fileops() {

	f, err := os.Create("log.txt")

	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = f.WriteString("hello welcome")
	if err != nil {
		log.Fatal(err.Error())
	}
	f.Close()
	if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Println("File exists.")
	}

	f, err = os.Open("log.txt")
	data, err := io.ReadAll(f)
	fmt.Println("", string(data))
	f.Close()
	f, err = os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(" \n appened to file ")
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	err = os.Remove("log.txt")
	if err != nil {
		log.Fatal(err)
	}

}

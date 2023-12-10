package main

import (
	"fmt"
	"os"
)

var m = map[string]string{

	"apple":  "olma",
	"age ":   "yosh",
	"digit ": "raqam",
	"cloud":  "bulut",
	"pen":    "ruchka",
	"pencil": "qalam",
	"hello":  "salom",
	"world":  "dunyo",
	"boys":   "bolalar",
	"girls":  "qizlar",
}

func main() {
	file, err := os.OpenFile("dictionary.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error", err)
	}
	defer file.Close()

	c := 0
	for i, v := range m {
		c++
		str := fmt.Sprintf("%d. %s - %s\n ", c, i, v)
		fmt.Println(str)
		_, err := file.Write([]byte(str))
		fmt.Println()
		if err != nil {
			panic(err)
		}
	}

}

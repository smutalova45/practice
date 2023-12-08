package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("errror is ", err.Error())
		return
	}

	defer file.Close()

	b, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("error *", err)
		return
	}
	var digit = []int{}
	var words = []string{}

	slice:=strings.Split(string(b)," ")

	for _, v := range slice {
		a, err := strconv.Atoi(v)
      if err == nil {
        digit = append(digit, a)
      } else {
        words = append(words, v)
      }
	}

	fmt.Println(words)
	fmt.Println(digit)

	fmt.Println(string(b))

}

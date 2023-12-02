package main

import (
	"fmt"
	"strings"
)

func main() {
	// a:=[]int32{12,23,45,0}
	// b:=[]int32{1,54,8,0}
	// fmt.Print(question(a,b))

	word := []string{"aziza", "abc", "ada"}
	fmt.Println(P(word))
	fmt.Println(interrupt("(al),(0)"))
}
func interrupt(comand string)string{

	result:=strings.Replace(comand,"()","0",-1)
	result=strings.Replace(result,"(al)","al",-1)
	return result
}
func P(word []string) string {
	for _, v := range word {
		if IsPalindrom(v) {
			return v
		}

	}
	return ""
}
func IsPalindrom(word string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}

	}
	return true
}
func question(a []int32, b []int32) []int32 {

	var counterA, counterB int32

	for i := 0; i < len(a); i++ {

		if a[i] > b[i] {
			counterA++
		} else if a[i] < b[i] {
			counterB++
		}

	}

	var result = [2]int32{counterA, counterB}

	return result[:]
}

package main

import (
	"fmt"
	"strings"
)

func wordNgram(str string, n int) []string {
	arr := strings.Split(str, " ")

	if len(arr) < n {
		return arr
	}

	var res []string
	for i := 0; i < len(arr)-1; i++ {
		res = append(res, strings.Join(arr[i:i+n], ""))
		fmt.Printf("%s", arr[i:i+n])
	}

	return res
}

func letterNgram(str string, n int) []string {
	arr := strings.Split(str, "")

	if len(arr) < n {
		return arr
	}

	var res []string
	for i := 0; i < len(arr)-1; i++ {
		res = append(res, strings.Join(arr[i:i+n], ""))
		fmt.Printf("%s", arr[i:i+n])
	}

	return res
}

func main() {
	text := "I am an NLPer"
	letterNgram(text, 2)
	fmt.Println("")
	wordNgram(text, 2)
}

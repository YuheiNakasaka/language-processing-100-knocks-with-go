package exercise

import "strings"

func exercise() map[string]int {
	element := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	arr := strings.Split(element, " ")
	var res = make(map[string]int)
	for i, v := range arr {
		i = i + 1
		if i == 1 || i == 5 || i == 6 || i == 7 || i == 8 || i == 9 || i == 15 || i == 16 || i == 19 {
			b := []byte(v)
			res[string(b[0])] = i
		} else {
			b := []byte(v)
			res[string(b[0])+string(b[1])] = i
		}
	}
	return res
}

package exercise

import "strings"

func exercise() string {
	str := strings.Split("パタトクカシーー", "")
	res := ""
	for i := 0; i < len(str)-1; i++ {
		if i%2 == 0 {
			res += str[i]
		}
	}
	return res
}

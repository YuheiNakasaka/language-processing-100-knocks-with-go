package exercise

import "strings"

func exercise() []int {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	arr := strings.Split(str, " ")
	res := []int{}
	for _, v := range arr {
		if strings.Contains(v, ",") {
			arrr := strings.Split(v, ",")
			for _, vv := range arrr {
				if len(vv) == 0 {
					res = append(res, 1)
				} else {
					res = append(res, len(vv))
				}
			}
		} else if strings.Contains(v, ".") {
			arrr := strings.Split(v, ".")
			for _, vv := range arrr {
				if len(vv) == 0 {
					res = append(res, 1)
				} else {
					res = append(res, len(vv))
				}
			}
		} else {
			res = append(res, len(v))
		}
	}
	return res
}

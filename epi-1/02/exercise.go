package exercise

import "strings"

func exercise() string {
	a := strings.Split("パトカー", "")
	b := strings.Split("タクシー", "")
	c := []string{}
	for i := 0; i < len(a); i++ {
		c = append(c, a[i])
		c = append(c, b[i])
	}
	return strings.Join(c, "")
}

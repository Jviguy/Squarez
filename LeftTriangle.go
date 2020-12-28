package Squarez

import "strings"

type LeftTriangle struct {
	rect *RECT
}

func (t *LeftTriangle) ToLeft() *LeftTriangle {
	return (*LeftTriangle)(t)
}

func (t *LeftTriangle) ToString() string {
	var str = strings.Repeat("_",t.rect.width)
	for i := 0; i < t.rect.width; i++ {
		str += Reverse("/"+strings.Repeat(" ",t.rect.width-i) + "|\n")
	}
	str += "\n-"
	return str
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
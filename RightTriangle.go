package Squarez

import "strings"

type RightTriangle struct {
	rect *RECT
}

func (t *RightTriangle) ToLeft() *LeftTriangle {
	return (*LeftTriangle)(t)
}

func (t *RightTriangle) ToString() string {
	var str = "-\n"
	for i := 0; i < t.rect.width-2; i++ {
		str += "|" + strings.Repeat(" ", i) + "\\\n"
	}
	str += strings.Repeat("_",t.rect.width)
	return str
}

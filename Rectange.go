package Squarez

import (
	"fmt"
	"io"
	"strings"
)

type RECT struct {
	width int
	height int
}

func (r *RECT) ExtendWidth(addedwith int)  {
	r.width += addedwith
}

func (r *RECT) SetWidth(newwitdh int)  {
	r.width = newwitdh
}

func (r *RECT) ShrinkWidth(subwidth int)  {
	r.width -= subwidth
}

func (r *RECT) GetWidth() int {
	return r.width
}

func (r *RECT) ShrinkHeight(subheight int)  {
	r.height -= subheight
}

func (r *RECT) ExtendHeight(addedheight int) {
	r.height += addedheight
}

func (r *RECT) SetHeight(newheight int)  {
	r.height = newheight
}

func (r *RECT) GetHeight() int {
	return r.height
}

func (r *RECT) PrintTo(w io.Writer) error {
	_, err := w.Write([]byte(r.ToString()))
	return err
}

func (r *RECT) PrintlnTo(w io.Writer) error {
	_, err := w.Write([]byte(r.ToString()+"\n"))
	return err
}

func (r *RECT) ToString() string {
	str := strings.Repeat("-",r.width)+"\n"
	str += strings.Repeat("|"+strings.Repeat(" ",r.width-2)+"|\n",r.height-2)
	str += strings.Repeat("-",r.width)+"\n"
	return str
}

func (r *RECT) FmtPrint() {
	fmt.Print(r.ToString())
}

//returns self for a builder like usage
func (r *RECT) Flip() *RECT {
	return NewRectangle(r.height,r.width)
}

func (r *RECT) FmtPrintln() {
	fmt.Println(r.ToString())
}

func (r *RECT) ToTriangles() (*RightTriangle,*LeftTriangle) {
	return &RightTriangle{r}, &LeftTriangle{r}
}

func NewRectangle(width int,height int) *RECT {
	return &RECT{width,height}
}
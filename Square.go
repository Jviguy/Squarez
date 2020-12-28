package Squarez

import (
	"fmt"
	"io"
	"strings"
)

type Square struct {
	size int
}

func (s *Square) ExtendBy(size int) {
	s.size += size
}

func (s *Square) Extend() {
	s.size++
}

func (s *Square) ShrinkBy(size int)  {
	s.size -= size
}

func (s *Square) Shrink()  {
	s.size--
}

func (s *Square) GetSize() int {
	return s.size
}

func (s *Square) SetSize(size int) {
	s.size = size
}

func (s *Square) PrintTo(w io.Writer) error {
	_, err := w.Write([]byte(s.ToString()))
	return err
}

func (s *Square) PrintlnTo(w io.Writer) error {
	_,err := w.Write([]byte(s.ToString()+"\n"))
	return err
}

func (s *Square) FmtPrint() {
	fmt.Print(s.ToString())
}

func (s *Square) FmtPrintln() {
	fmt.Println(s.ToString())
}

func (s *Square) ToString() string {
	str := strings.Repeat("-",s.size)+"\n"
	str += strings.Repeat("|"+strings.Repeat(" ",s.size-2)+"|\n",s.size-2)
	str += strings.Repeat("-",s.size)+"\n"
	return str
}

func (s *Square) ToTriangles() (*RightTriangle,*LeftTriangle) {
	return &RightTriangle{rect: NewRectangle(s.size,s.size)},&LeftTriangle{rect: NewRectangle(s.size,s.size)}
}

func NewSquare(size int) *Square {
	return &Square{size}
}


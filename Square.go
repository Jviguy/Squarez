/**

        :::   :::       :::     :::::::::  ::::::::::
      :+:+: :+:+:    :+: :+:   :+:    :+: :+:
    +:+ +:+:+ +:+  +:+   +:+  +:+    +:+ +:+
   +#+  +:+  +#+ +#++:++#++: +#+    +:+ +#++:++#
  +#+       +#+ +#+     +#+ +#+    +#+ +#+
 #+#       #+# #+#     #+# #+#    #+# #+#
###       ### ###     ### #########  ##########

            :::::::::  :::   :::
            :+:    :+: :+:   :+:
            +:+    +:+  +:+ +:+
            +#++:++#+    +#++:
            +#+    +#+    +#+
            #+#    #+#    #+#
            #########     ###

     ::::::::::: :::     ::: ::::::::::: ::::::::  :::    ::: :::   :::
        :+:     :+:     :+:     :+:    :+:    :+: :+:    :+: :+:   :+:
       +:+     +:+     +:+     +:+    +:+        +:+    +:+  +:+ +:+
      +#+     +#+     +:+     +#+    :#:        +#+    +:+   +#++:
     +#+      +#+   +#+      +#+    +#+   +#+# +#+    +#+    +#+
#+# #+#       #+#+#+#       #+#    #+#    #+# #+#    #+#    #+#
 #####          ###     ########### ########   ########     ###

MIT License

Copyright (c) 2020 Jviguy

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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


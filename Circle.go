package Squarez

import (
	"math"
	"strings"
)

type Circle struct {
	circumference int
}

func (c *Circle) ToString() string {
	area := c.calculateArea()
	str := " " + strings.Repeat("-",area/2)+"\n"
	str += strings.Repeat("|" +strings.Repeat(" ",area/2)+"|\n",area/2)
	str += " " + strings.Repeat("-",area/2)+"\n"
	return str
}

//This function tells us how much area we have to leave in the middle
func (c *Circle) calculateArea() int {
	return int(math.RoundToEven(math.Pow(float64(c.circumference),2)/(4*math.Pi)))
}

func NewCircle(circumference int) *Circle {
	return &Circle{circumference}
}

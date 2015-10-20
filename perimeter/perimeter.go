package shape

import "math"

type Shape interface {
  area() float64
  perimeter() float64
}

type Circle struct {
  radius float64
}

type Rectangle struct {
  length float64
  breadth float64
}


func (circ Circle)  perimeter() float64 {
  return 2 * math.Pi * circ.radius
}

func (circ Circle) area() float64 {
  return math.Pi * circ.radius * circ.radius
}

func (rect Rectangle) perimeter() float64 {
  return 2 * (rect.length + rect.breadth)
}

func (rect Rectangle) area() float64 {
  return rect.length * rect.breadth
}

func GetArea(s Shape) float64{
    return s.area()
}

func GetPerimeter(s Shape) float64{
    return s.perimeter()
}

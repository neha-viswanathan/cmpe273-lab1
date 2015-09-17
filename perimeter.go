package main

import ("fmt";"math")

type Circle struct {
  radius float64
}

type Rectangle struct {
  length float64
  breadth float64
}

func PerimeterCircle (circ Circle) float64 {
  return 2 * math.Pi * circ.radius
}

func AreaCircle (circ Circle) float64 {
  return math.Pi * circ.radius * circ.radius
}

func PerimeterRectangle (rect Rectangle) float64 {
  return 2 * (rect.length + rect.breadth)
}

func AreaRectangle (rect Rectangle) float64 {
  return rect.length * rect.breadth
}

func main () {
  circ := Circle{1}

  fmt.Println("Enter radius of circle")
  fmt.Scanln(&circ.radius)
  if(circ.radius <= 0) {
    fmt.Println("Invalid/Incorrect value for radius. It cannot be zero or negative!")
  } else {
    fmt.Println("The perimeter of circle is", PerimeterCircle(circ), "units")
    fmt.Println("The area of circle is", AreaCircle(circ), "units")
  }

  rect := Rectangle{3, 2}
  fmt.Println("Enter length of rectangle")
  fmt.Scanln(&rect.length)
  fmt.Println("Enter breadth of rectangle")
  fmt.Scanln(&rect.breadth)
  if(rect.length <= 0 || rect.breadth <= 0) {
    fmt.Println("Invalid/Incorrect value for length/breadth. It cannot be zero or negative!")
  } else {
    fmt.Println("The perimeter of rectangle is", PerimeterRectangle(rect), "units")
    fmt.Println("The area of rectangle is", AreaRectangle(rect), "units")
  }
}

package main

import "math"

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}


func (t Triangle) Area() float64 {
  return 0 
}

func (t Triangle) Perimeter() float64 {
 return (t.Base + t.Height) * 2
}


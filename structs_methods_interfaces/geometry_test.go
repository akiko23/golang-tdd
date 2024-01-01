package main

import (
	"testing"
)

type TestShape struct {
  name string
  shape Shape
  want float64
}

func TestPerimeter(t *testing.T) {
  assertPerimeter := func(t testing.TB, shape Shape, want float64) {
    t.Helper()
    got := shape.Perimeter()
    if got != want {
      t.Errorf("Expected %g got %g", want, got)
    }
  }

  perimeterTests := []TestShape {
    {name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 40.0},
    {name: "Circle", shape: Circle{Radius: 10.0}, want: 62.83185307179586},
    {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
  }
  
  for _, tt := range perimeterTests {
    t.Run(tt.name, func(t *testing.T) {
      assertPerimeter(t, tt.shape, tt.want)
    })
  }
}

func TestArea(t *testing.T) {
  
  assertArea := func(t testing.TB, shape Shape, want float64) {
    t.Helper()
    got := shape.Area()
    
    if got != want {
      t.Errorf("Expected %g but got %g", want, got)
    }
  }

  areaTests := []TestShape {
    {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
    {name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793}, 
  }
  
  for _, tt := range areaTests {
    t.Run(tt.name, func(t *testing.T) {
      assertArea(t, tt.shape, tt.want)
    })
  }

}

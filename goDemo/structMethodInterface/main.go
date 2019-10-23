package main

import "math"

// Shape 接口
type Shape interface {
	Area() float64
}

// Rectangle 长方形
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle 圆形
type Circle struct {
	Radius float64
}

// Triangle 三角形
type Triangle struct {
	Base   float64
	Height float64
}

// Area 长方形面积
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area 圆形面积
func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Area 三角形面积
func (t *Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

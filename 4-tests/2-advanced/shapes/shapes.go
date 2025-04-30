package shapes

// Shape is an interface that defines a method for calculating the area of a shape.
// It can be implemented by any type that has an area method.
type Shape interface {
	area() float64
}

// Rectangle and Circle are concrete types that implement the Shape interface.
// They have their own area methods that calculate the area of the respective shapes.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for Rectangle calculates the area of the rectangle using the formula width * height.
// The method is defined on the Rectangle type and implements the Shape interface.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a concrete type that implements the Shape interface.
// It has its own area method that calculates the area of a circle.
type Circle struct {
	Radius float64
}

// Area method for Circle calculates the area of the circle using the formula πr².
// The value of π is approximated to 3.14 for simplicity.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

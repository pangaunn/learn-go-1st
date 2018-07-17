package aaa

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type geometry interface {
	area() float64
	peremeter() float64
}

func (r Rectangle) peremeter() float64 {
	return r.width * r.height
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) peremeter() float64 {
	return 2 * 3.1416 * c.radius
}

func (c Circle) area() float64 {
	return 3.1416 * c.radius * c.radius
}

func Peremeter(g geometry) float64 {
	// return 2 * (r.width + r.height)
	return g.peremeter()
}

func Area(g geometry) float64 {
	// return r.width * r.height
	return g.area()
}

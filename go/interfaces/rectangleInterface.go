package interfaces

type Shape interface {
	Area() float64
	Perimeter() float64
}


type Rectangle struct{
	Length, Width float64
}


func (r Rectangle) Area() float64{
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64{
	return 2*( r.Length + r.Width)
}
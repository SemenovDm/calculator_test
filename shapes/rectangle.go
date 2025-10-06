package shapes

type Rectangle struct {
	Width float64
	Heigh float64
}

func (r Rectangle) Area() float64 {

	return 2 * (r.Heigh + r.Width)

}

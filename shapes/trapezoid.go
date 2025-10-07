package shapes

type Trapezoid struct {
	length1 float64
	length2 float64
	base    float64
}

func (t Trapezoid) Area() float64 {

	return 0.5 * (t.length1 + t.length2) * t.base

}

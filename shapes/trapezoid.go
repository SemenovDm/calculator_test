package shapes

type Trapezoid struct {
	Length1 float64
	Length2 float64
	Base    float64
}

func (t Trapezoid) Area() float64 {

	return 0.5 * (t.Length1 + t.Length2) * t.Base

}

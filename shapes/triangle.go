package shapes

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {

	return 0.5 * t.Base * t.Height

}

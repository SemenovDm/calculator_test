package calculator

import "calculator_test/shapes"

func TotalArea(circles []shapes.Circle, trianles []shapes.Triangle, rectanles []shapes.Rectangle) float64 {

	ta := 0.0

	for _, c := range circles {

		ta += c.Area()
	}

	for _, t := range trianles {
		ta += t.Area()
	}

	for _, r := range rectanles {
		ta += r.Area()
	}

	return ta

}

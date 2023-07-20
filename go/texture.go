package main

import "math"

type texture interface {
	value(u float64, v float64, p point3) color
}

type solidColor struct {
	colorValue color
}

func newSolidColor(col color) solidColor {
	return solidColor{colorValue: col}
}

func (c solidColor) value(u float64, v float64, p point3) color {
	return c.colorValue
}

type checkerTexture struct {
	odd  texture
	even texture
}

func newSolidChecker(col0 color, col1 color) checkerTexture {
	return checkerTexture{odd: newSolidColor(col0), even: newSolidColor(col1)}
}

func newChecker(t0 texture, t1 texture) checkerTexture {
	return checkerTexture{odd: t0, even: t1}
}

func (c checkerTexture) value(u float64, v float64, p point3) color {
	sines := math.Sin(10*p.x()) * math.Sin(10*p.y()) * math.Sin(10*p.z())
	if sines < 0 {
		return c.odd.value(u, v, p)
	}
	return c.even.value(u, v, p)
}

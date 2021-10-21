package main

import (
	"fmt"
	"math"
)

type color = vec3

func newColor(e0 float64, e1 float64, e2 float64) color {
	return color{e: [3]float64{e0, e1, e2}}
}

func writeColor(pixelColor color, samplesPerPixel int) {
	r := pixelColor.x()
	g := pixelColor.y()
	b := pixelColor.z()

	// Divide by number of samples and correct for gamma=2.0
	scale := 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(r * scale)
	g = math.Sqrt(g * scale)
	b = math.Sqrt(b * scale)

	fmt.Println(int(255.999*clamp(r, 0.0, 0.999)), " ", int(255.999*clamp(g, 0.0, 0.999)), " ", int(255.999*clamp(b, 0.0, 0.999)))
}

package main

import (
	"math"
	"math/rand"
)

type camera struct {
	origin               point3
	horizontal, vertical vec3
	lowerLeftCorner      point3
	lensRadius           float64
	u, v                 vec3
	time0, time1         float64
}

func newCamera(lookFrom point3, lookAt point3, vup vec3, vfov float64, aspectRatio float64, aperture float64, focusDist float64, t0 float64, t1 float64) camera {
	c := camera{}

	theta := degToRad(vfov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := unitVector(sub(lookFrom, lookAt))
	c.u = unitVector(cross(vup, w))
	c.v = cross(w, c.u)

	c.origin = lookFrom
	c.horizontal = mult(focusDist*viewportWidth, c.u)
	c.vertical = mult(focusDist*viewportHeight, c.v)
	// origin - horizontal/2 - vertical/2 - w
	c.lowerLeftCorner = sub(sub(sub(c.origin, div(c.horizontal, 2)), div(c.vertical, 2)), mult(focusDist, w))

	c.lensRadius = aperture / 2.0
	c.time0 = t0
	c.time1 = t1

	return c
}

func (c camera) getRay(s float64, t float64, generator *rand.Rand) ray {
	rd := mult(c.lensRadius, randomInDisk(generator))
	offset := add(mult(rd.x(), c.u), mult(rd.y(), c.v))
	// Dir = lowerLeftCorner + s*horizontal + t*vertical - origin
	return newRay(add(c.origin, offset), sub(sub(add(add(c.lowerLeftCorner, mult(s, c.horizontal)), mult(t, c.vertical)), c.origin), offset), randFloat(c.time0, c.time1, generator))
}

package main

import (
	"fmt"
	"math"
	"math/rand"
)

type vec3 struct {
	e [3]float64
}

func newVec3(e0 float64, e1 float64, e2 float64) vec3 {
	return vec3{e: [3]float64{e0, e1, e2}}
}

func (v vec3) x() float64 {
	return v.e[0]
}

func (v vec3) y() float64 {
	return v.e[1]
}

func (v vec3) z() float64 {
	return v.e[2]
}

func (v vec3) neg() vec3 {
	u := vec3{e: [3]float64{-v.e[0], -v.e[1], -v.e[2]}}
	return u
}

func (v *vec3) add(u vec3) {
	v.e[0] += u.e[0]
	v.e[1] += u.e[1]
	v.e[2] += u.e[2]
}

func (v *vec3) mult(t float64) {
	v.e[0] *= t
	v.e[1] *= t
	v.e[2] *= t
}

func (v *vec3) div(t float64) {
	v.mult(1 / t)
}

func (v vec3) magSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v vec3) mag() float64 {
	return math.Sqrt(v.magSquared())
}

func (v vec3) out() {
	fmt.Println(v.e[0], " ", v.e[1], " ", v.e[2])
}

type point3 = vec3

func newPoint3(e0 float64, e1 float64, e2 float64) point3 {
	return point3{e: [3]float64{e0, e1, e2}}
}

// Utility functions
func add(v vec3, u vec3) vec3 {
	return vec3{e: [3]float64{v.e[0] + u.e[0], v.e[1] + u.e[1], v.e[2] + u.e[2]}}
}

// Returns v - u
func sub(v vec3, u vec3) vec3 {
	return vec3{e: [3]float64{v.e[0] - u.e[0], v.e[1] - u.e[1], v.e[2] - u.e[2]}}
}

func prod(v vec3, u vec3) vec3 {
	return vec3{e: [3]float64{v.e[0] * u.e[0], v.e[1] * u.e[1], v.e[2] * u.e[2]}}
}

func mult(t float64, v vec3) vec3 {
	return vec3{e: [3]float64{v.e[0] * t, v.e[1] * t, v.e[2] * t}}
}

func div(v vec3, t float64) vec3 {
	return vec3{e: [3]float64{v.e[0] / t, v.e[1] / t, v.e[2] / t}}
}

func dot(v vec3, u vec3) float64 {
	return v.e[0]*u.e[0] + v.e[1]*u.e[1] + v.e[2]*u.e[2]
}

func cross(v vec3, u vec3) vec3 {
	return vec3{e: [3]float64{v.e[1]*u.e[2] - v.e[2]*u.e[1], v.e[2]*u.e[0] - v.e[0]*u.e[2], v.e[0]*u.e[1] - v.e[1]*u.e[0]}}
}

func unitVector(v vec3) vec3 {
	return div(v, v.mag())
}

func reflect(v vec3, n vec3) vec3 {
	// v - 2*dot(v,n)*n
	return sub(v, mult(2*dot(v, n), n))
}

func refract(v vec3, n vec3, refraction float64) vec3 {
	cosTheta := min(dot(v.neg(), n), 1.0)
	rOutPerp := mult(refraction, add(v, mult(cosTheta, n)))
	rOutPar := mult(-1.0*math.Sqrt(math.Abs(1-rOutPerp.magSquared())), n)
	return add(rOutPar, rOutPerp)
}

func (v vec3) nearZero() bool {
	s := 1e-8
	return (math.Abs(v.e[0]) < s) && (math.Abs(v.e[1]) < s) && (math.Abs(v.e[2]) < s)
}

// Random functions
func random(generator *rand.Rand) vec3 {
	return vec3{e: [3]float64{generator.Float64(), generator.Float64(), generator.Float64()}}
}

func randomLim(min float64, max float64, generator *rand.Rand) vec3 {
	return vec3{e: [3]float64{randFloat(min, max, generator), randFloat(min, max, generator), randFloat(min, max, generator)}}
}

func randomInSphere(generator *rand.Rand) vec3 {
	for {
		p := randomLim(-1, 1, generator)
		if p.magSquared() < 1 {
			return p
		}
	}
}

func randomInDisk(generator *rand.Rand) vec3 {
	for {
		p := vec3{e: [3]float64{randFloat(-1.0, 1.0, generator), randFloat(-1.0, 1.0, generator), 0}}
		if p.magSquared() < 1 {
			return p
		}
	}
}

func randomInHemisphere(normal vec3, generator *rand.Rand) vec3 {
	p := randomInSphere(generator)
	if dot(p, normal) > 0.0 {
		return p
	}
	return p.neg()
}

func randomUnitVector(generator *rand.Rand) vec3 {
	return unitVector(randomInSphere(generator))
}

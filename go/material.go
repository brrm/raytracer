package main

import (
	"math"
	"math/rand"
)

type material interface {
	scatter(rIn ray, rec *hitRecord, attenuation *color, scattered *ray, generator *rand.Rand) bool
}

// Lambertian - matte
type lambertian struct {
	albedo texture
}

func newLambertianSolid(a color) lambertian {
	return lambertian{albedo: newSolidColor(a)}
}

func newLambertian(a texture) lambertian {
	return lambertian{albedo: a}
}

func (l lambertian) scatter(rIn ray, rec *hitRecord, attenuation *color, scattered *ray, generator *rand.Rand) bool {
	scatterDirection := add(rec.normal, randomUnitVector(generator))
	if scatterDirection.nearZero() {
		scatterDirection = rec.normal
	}
	*scattered = newRay(rec.p, scatterDirection, rIn.time())
	*attenuation = l.albedo.value(rec.u, rec.v, rec.p)
	return true
}

// Metal - reflects
type metal struct {
	albedo color
	fuzz   float64
}

func newMetal(a color, f float64) metal {
	return metal{albedo: a, fuzz: min(f, 1.0)}
}

func (m metal) scatter(rIn ray, rec *hitRecord, attenuation *color, scattered *ray, generator *rand.Rand) bool {
	reflected := reflect(unitVector(rIn.direction()), rec.normal)
	*scattered = newRay(rec.p, add(reflected, mult(m.fuzz, randomUnitVector(generator))), rIn.time())
	*attenuation = m.albedo
	return dot(scattered.direction(), rec.normal) > 0
}

// Dielectric - refracts
type dielectric struct {
	ir float64
}

func newDielectric(refractiveIndex float64) dielectric {
	return dielectric{ir: refractiveIndex}
}

func (d dielectric) scatter(rIn ray, rec *hitRecord, attenuation *color, scattered *ray, generator *rand.Rand) bool {
	*attenuation = newColor(1.0, 1.0, 1.0)
	var refractionRatio float64
	if rec.frontFace {
		refractionRatio = 1.0 / d.ir
	} else {
		refractionRatio = d.ir
	}

	unitDirection := unitVector(rIn.direction())
	cosTheta := min(dot(unitDirection.neg(), rec.normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	direction := vec3{}
	cannotRefract := sinTheta*refractionRatio > 1.0
	// Reflect
	if cannotRefract || schlick(cosTheta, refractionRatio) > generator.Float64() {
		direction = reflect(unitDirection, rec.normal)
	} else {
		direction = refract(unitDirection, rec.normal, refractionRatio)
	}

	*scattered = newRay(rec.p, direction, rIn.time())
	return true
}

func schlick(cosTheta float64, refractiveIndex float64) float64 {
	r0 := (1.0 - refractiveIndex) / (1.0 + refractiveIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosTheta, 5)
}

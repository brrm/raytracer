package main

import (
	"math/rand"
	"time"
)

func randomScene() hittableList {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	world := hittableList{}
	//groundMaterial := newLambertian(newSolidChecker(newColor(0.2, 0.3, 0.1), newColor(0.9, 0.9, 0.9)))
	groundMaterial := newLambertianSolid(newColor(0.5, 0.5, 0.5))
	world.add(newSphere(newPoint3(0, -1000, 0), 1000, groundMaterial))
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := generator.Float64()
			center := newPoint3(float64(a)+0.9*generator.Float64(), 0.2, float64(b)+0.9*generator.Float64())

			if sub(center, newPoint3(4, 0.2, 0)).mag() > 0.9 {
				var sphereMaterial material

				if chooseMat < 0.8 {
					// Diffuse
					albedo := prod(random(generator), random(generator))
					sphereMaterial = newLambertianSolid(albedo)
					world.add(newSphere(center, 0.2, sphereMaterial))
					center2 := add(center, newPoint3(0, randFloat(0, 0.5, generator), 0))
					world.add(newMovingSphere(center, center2, 0.0, 1.0, 0.2, sphereMaterial))
				} else if chooseMat < 0.95 {
					// Metal
					albedo := randomLim(0.5, 1, generator)
					fuzz := randFloat(0, 0.5, generator)
					sphereMaterial = newMetal(albedo, fuzz)
					world.add(newSphere(center, 0.2, sphereMaterial))
				} else {
					// Dielectric
					sphereMaterial = newDielectric(1.5)
					world.add(newSphere(center, 0.2, sphereMaterial))
				}
			}
		}
	}

	materialOne := newDielectric(1.5)
	world.add(newSphere(newPoint3(0, 1, 0), 1, materialOne))
	materialTwo := newLambertianSolid(newColor(0.4, 0.2, 0.1))
	world.add(newSphere(newPoint3(-4, 1, 0), 1, materialTwo))
	materialThree := newMetal(newColor(0.7, 0.6, 0.5), 0.0)
	world.add(newSphere(newPoint3(4, 1, 0), 1, materialThree))

	return hittableList{objects: []hittable{newBvh(world.objects, 0, len(world.objects), 0, 0.0)}}
}

func simpleScene() hittableList {
	world := hittableList{}

	materialGround := newLambertianSolid(newColor(0.8, 0.8, 0))
	materialCenter := newLambertianSolid(newColor(0.1, 0.2, 0.5))
	materialLeft := newDielectric(1.5)
	materialRight := newMetal(newColor(0.8, 0.6, 0.2), 0)

	world.add(newSphere(newPoint3(0, -100.5, -1), 100, materialGround))
	world.add(newSphere(newPoint3(0, 0, -1), 0.5, materialCenter))
	world.add(newSphere(newPoint3(-1, 0, -1), 0.5, materialLeft))
	world.add(newSphere(newPoint3(-1, 0, -1), -0.45, materialLeft))
	world.add(newSphere(newPoint3(1, 0, -1), 0.5, materialRight))

	return hittableList{objects: []hittable{newBvh(world.objects, 0, len(world.objects), 0, 1)}}
}

func twoSpheres() hittableList {
	world := hittableList{}
	checker := newLambertian(newSolidChecker(newColor(0.2, 0.3, 0.1), newColor(0.9, 0.9, 0.9)))
	world.add(newSphere(newPoint3(0, -10, 0), 10, checker))
	world.add(newSphere(newPoint3(0, 10, 0), 10, checker))
	return world
}

func demoScene() hittableList {

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	world := hittableList{}
	groundMaterial := newLambertian(newSolidChecker(newColor(0.1, 0.1, 0.1), newColor(0.9, 0.9, 0.9)))
	world.add(newSphere(newPoint3(0, -1000, 0), 1000, groundMaterial))
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := generator.Float64()
			center := newPoint3(float64(a)+0.9*generator.Float64(), 0.2, float64(b)+0.9*generator.Float64())

			if sub(center, newPoint3(4, 0.2, 0)).mag() > 0.9 {
				var sphereMaterial material

				if chooseMat < 0.75 {
					// Diffuse
					albedo := prod(random(generator), random(generator))
					sphereMaterial = newLambertianSolid(albedo)
					world.add(newSphere(center, 0.2, sphereMaterial))
					center2 := add(center, newPoint3(0, randFloat(0, 0.5, generator), 0))
					world.add(newMovingSphere(center, center2, 0.0, 1.0, 0.2, sphereMaterial))
				} else if chooseMat < 0.90 {
					// Metal
					albedo := randomLim(0.5, 1, generator)
					fuzz := randFloat(0, 0.5, generator)
					sphereMaterial = newMetal(albedo, fuzz)
					world.add(newSphere(center, 0.2, sphereMaterial))
				} else {
					// Dielectric
					sphereMaterial = newDielectric(1.5)
					world.add(newSphere(center, 0.2, sphereMaterial))
				}
			}
		}
	}

	materialGlass := newDielectric(1.5)
	world.add(newSphere(newPoint3(0, 1, 0), 1, materialGlass))
	materialMetal := newMetal(newColor(0.7, 0.6, 0.5), 0.0)
	world.add(newSphere(newPoint3(4, 1, 0), 1, materialMetal))

	return hittableList{objects: []hittable{newBvh(world.objects, 0, len(world.objects), 0, 0.0)}}
}

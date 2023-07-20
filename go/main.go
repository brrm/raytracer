package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"
)

func rayColor(r ray, w hittableList, depth int, generator *rand.Rand) color {
	totalAttenuation := newColor(1, 1, 1)
	attenuation := color{}
	rec := hitRecord{}
	for i := 0; i < depth; i++ {
		if w.hit(r, 0.001, math.Inf(1), &rec) {
			if (*rec.matPtr).scatter(r, &rec, &attenuation, &r, generator) {
				totalAttenuation = prod(attenuation, totalAttenuation)
			} else {
				return newColor(0, 0, 0)
			}
		} else {
			unitDirection := unitVector(r.direction())
			t := 0.5 * (unitDirection.y() + 1.0)
			return prod(totalAttenuation, add(mult((1.0-t), newColor(1.0, 1.0, 1.0)), mult(t, newColor(0.5, 0.7, 1.0))))
		}
	}
	return newColor(0, 0, 0)
}

func main() {
	// Image
	const aspectRatio = 16.0 / 9.0
	const imageWidth = 1600
	const imageHeight = int(float64(imageWidth) / aspectRatio)
	const samplesPerPixel = 50
	const maxDepth = 50

	// World
	var world hittableList
	var lookFrom, lookAt point3
	vfov := 40.0
	aperture := 0.0
	switch 3 {
	case 0:
		world = randomScene()
		lookFrom = newPoint3(13, 2, 3)
		lookAt = newPoint3(0, 0, 0)
		vfov = 20.0
		aperture = 0.1
	case 1:
		world = simpleScene()
		lookFrom = newPoint3(-2, 2, 1)
		lookAt = newPoint3(0, 0, -1)
		vfov = 20.0
		aperture = 0.1
	case 2:
		world = twoSpheres()
		lookFrom = newPoint3(13, 2, 3)
		lookAt = newPoint3(0, 0, 0)
		vfov = 20.0
	case 3:
		world = demoScene()
		lookFrom = newPoint3(13, 2, 3)
		lookAt = newPoint3(0, 0, 0)
		vfov = 20.0
		aperture = 0.1
	}

	// Camera
	vup := newVec3(0, 1, 0)
	focalDist := 10.0
	camera := newCamera(lookFrom, lookAt, vup, vfov, aspectRatio, aperture, focalDist, 0.0, 0.0)

	// Render
	numThreads := 8
	start := time.Now()

	var img [imageWidth][imageHeight]color
	chunks := split(imageHeight, numThreads)

	var wg sync.WaitGroup
	wg.Add(numThreads)
	remaining := numThreads

	for t := 0; t < numThreads; t++ {
		go func(start int, stop int) {
			defer wg.Done()
			generator := rand.New(rand.NewSource(time.Now().UnixNano()))
			for j := start - 1; j >= stop; j-- {
				for i := 0; i < imageWidth; i++ {
					pixelColor := color{}
					for s := 0; s < samplesPerPixel; s++ {
						u := (float64(i) + generator.Float64()) / float64(imageWidth-1)
						v := (float64(j) + generator.Float64()) / float64(imageHeight-1)
						r := camera.getRay(u, v, generator)
						pixelColor.add(rayColor(r, world, maxDepth, generator))
					}
					img[i][j] = pixelColor
				}
			}
			remaining--
			fmt.Fprintf(os.Stderr, "\r%v threads remaining", remaining)
		}(chunks[t], chunks[t+1])
	}

	fmt.Fprintf(os.Stderr, "Waiting on threads...\n")
	wg.Wait()

	fmt.Println("P3\n", imageWidth, " ", imageHeight, "\n255")
	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			writeColor(img[i][j], samplesPerPixel)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone! in %v", time.Since(start))
}

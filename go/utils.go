package main

import "math/rand"

const pi float64 = 3.1415926535897932385

func randFloat(min float64, max float64, generator *rand.Rand) float64 {
	return min + (max-min)*generator.Float64()
}

func degToRad(degrees float64) float64 {
	return (degrees * pi) / 180.0
}

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func min(x float64, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func max(x float64, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

// For splitting for loops into n chunks
func split(size int, n int) []int {
	var cutoff = size / n
	chunks := []int{size, size - cutoff}
	for {
		chunk := chunks[len(chunks)-1] - cutoff
		if chunk < cutoff {
			chunks = append(chunks, 0)
			return chunks
		}
		chunks = append(chunks, chunk)
	}
}

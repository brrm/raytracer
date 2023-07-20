package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type bvhNode struct {
	box   aabb
	left  hittable
	right hittable
}

func (b bvhNode) boundingBox(time0 float64, time1 float64, box *aabb) bool {
	*box = b.box
	return true
}

func (b bvhNode) hit(r ray, tMin float64, tMax float64, rec *hitRecord) bool {
	if !b.box.hit(r, tMin, tMax, rec) {
		return false
	}
	hitLeft := b.left.hit(r, tMin, tMax, rec)
	var hitRight bool
	if hitLeft {
		hitRight = b.right.hit(r, tMin, (*rec).t, rec)
	} else {
		hitRight = b.right.hit(r, tMin, tMax, rec)
	}
	return hitLeft || hitRight
}

func newBvh(objects []hittable, start int, end int, time0 float64, time1 float64) bvhNode {
	node := bvhNode{}

	axis := rand.Intn(3)
	span := end - start
	if span == 1 {
		node.left = objects[start]
		node.right = objects[start]
	} else if span == 2 {
		if boxCompare(objects[start], objects[start+1], axis) {
			node.left = objects[start]
			node.right = objects[start+1]
		} else {
			node.left = objects[start+1]
			node.right = objects[start]
		}
	} else {
		sort.Slice(objects[start:end], func(i, j int) bool { return boxCompare(objects[i], objects[j], axis) })
		var mid int = start + span/2
		node.left = newBvh(objects, start, mid, time0, time1)
		node.right = newBvh(objects, mid, end, time0, time1)
	}

	var boxLeft, boxRight aabb
	if !node.left.boundingBox(time0, time1, &boxLeft) || !node.right.boundingBox(time0, time1, &boxRight) {
		fmt.Fprintf(os.Stderr, "No bounding box when constructing bvh.\n")
	}
	node.box = surroundingBox(boxLeft, boxRight)

	return node
}

func boxCompare(a hittable, b hittable, axis int) bool {
	var boxA, boxB aabb
	if !a.boundingBox(0, 0, &boxA) || !b.boundingBox(0, 0, &boxB) {
		fmt.Fprintf(os.Stderr, "No bounding box when constructing bvh.\n")
	}
	return boxA.min().e[axis] < boxB.min().e[axis]
}

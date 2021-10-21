package main

import "math"

type movingSphere struct {
	cen0, cen1   point3
	time0, time1 float64
	radius       float64
	matPtr       *material
}

func newMovingSphere(c0 point3, c1 point3, t0 float64, t1 float64, r float64, mat material) movingSphere {
	return movingSphere{cen0: c0, cen1: c1, time0: t0, time1: t1, radius: r, matPtr: &mat}
}

func (s movingSphere) center(time float64) point3 {
	return add(s.cen0, mult((time-s.time0)/(s.time1-s.time0), sub(s.cen1, s.cen0)))
}

func (s movingSphere) hit(r ray, tMin float64, tMax float64, rec *hitRecord) bool {
	oc := sub(r.origin(), s.center(r.time()))
	a := r.dir.magSquared()
	halfB := dot(oc, r.direction())
	c := oc.magSquared() - s.radius*s.radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtD := math.Sqrt(discriminant)

	// Find nearest root that lies in the acceptable range
	root := (-halfB - sqrtD) / a
	if root < tMin || root > tMax {
		root = (-halfB + sqrtD) / a
		if root < tMin || root > tMax {
			return false
		}
	}

	rec.t = root
	rec.p = r.at(rec.t)
	// Dividing by radius yields unit vector
	outwardNormal := div(sub(rec.p, s.center(r.time())), s.radius)
	rec.setFaceNormal(r, outwardNormal)
	rec.matPtr = s.matPtr

	return true
}

func (s movingSphere) boundingBox(time0 float64, time1 float64, box *aabb) bool {
	box0 := newAabb(sub(s.center(time0), newVec3(s.radius, s.radius, s.radius)), add(s.center(time0), newVec3(s.radius, s.radius, s.radius)))
	box1 := newAabb(sub(s.center(time1), newVec3(s.radius, s.radius, s.radius)), add(s.center(time1), newVec3(s.radius, s.radius, s.radius)))
	*box = surroundingBox(box0, box1)
	return true
}

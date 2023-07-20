package main

import (
	"math"
)

type sphere struct {
	center point3
	radius float64
	matPtr *material
}

func newSphere(c point3, r float64, m material) sphere {
	return sphere{center: c, radius: r, matPtr: &m}
}

func (s sphere) hit(r ray, tMin float64, tMax float64, rec *hitRecord) bool {
	oc := sub(r.origin(), s.center)
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
	outwardNormal := div(sub(rec.p, s.center), s.radius)
	rec.setFaceNormal(r, outwardNormal)
	s.getUV(outwardNormal, &rec.u, &rec.v)
	rec.matPtr = s.matPtr

	return true
}

func (s sphere) boundingBox(time0 float64, time1 float64, box *aabb) bool {
	*box = newAabb(sub(s.center, newVec3(s.radius, s.radius, s.radius)), add(s.center, newVec3(s.radius, s.radius, s.radius)))
	return true
}

func (s sphere) getUV(p point3, u *float64, v *float64) {
	theta := math.Acos(-p.y())
	phi := math.Atan2(-p.z(), p.x()) + pi
	*u = phi / (2 * pi)
	*v = theta / pi
}

package main

type aabb struct {
	minimum point3
	maximum point3
}

func newAabb(a point3, b point3) aabb {
	return aabb{minimum: a, maximum: b}
}

func (a aabb) min() point3 {
	return a.minimum
}

func (a aabb) max() point3 {
	return a.maximum
}

func (a aabb) hit(r ray, tMin float64, tMax float64, rec *hitRecord) bool {
	// For each dimension
	for i := 0; i < 3; i++ {
		invD := 1.0 / r.direction().e[i]
		t0 := (a.min().e[i] - r.origin().e[i]) * invD
		t1 := (a.max().e[i] - r.origin().e[i]) * invD
		if invD < 0.0 {
			t0, t1 = t1, t0
		}

		tMin = max(t0, tMin)
		tMax = min(t1, tMax)

		if tMax <= tMin {
			return false
		}
	}
	return true
}

func surroundingBox(box0 aabb, box1 aabb) aabb {
	small := newPoint3(min(box0.min().x(), box1.min().x()), min(box0.min().y(), box1.min().y()), min(box0.min().z(), box1.min().z()))
	big := newPoint3(max(box0.max().x(), box1.max().x()), max(box0.max().y(), box1.max().y()), max(box0.max().z(), box1.max().z()))
	return newAabb(small, big)
}

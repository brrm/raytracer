package main

type hitRecord struct {
	p         point3
	normal    vec3
	matPtr    *material
	t         float64
	u, v      float64
	frontFace bool
}

func (rec *hitRecord) setFaceNormal(r ray, outwardNormal vec3) {
	// If true, then the ray hit from outside the sphere
	rec.frontFace = dot(r.direction(), outwardNormal) < 0
	if rec.frontFace {
		rec.normal = outwardNormal
	} else {
		rec.normal = outwardNormal.neg()
	}
}

type hittable interface {
	hit(r ray, tMin float64, tMax float64, rec *hitRecord) bool
	boundingBox(time0 float64, time1 float64, box *aabb) bool
}

type hittableList struct {
	objects []hittable
}

func (l *hittableList) clear() {
	l.objects = []hittable{}
}

func (l *hittableList) add(h hittable) {
	l.objects = append(l.objects, h)
}

func (l hittableList) hit(r ray, tMin float64, tMax float64, rec *hitRecord) bool {
	tempRec := hitRecord{}
	hitAnything := false
	closest := tMax

	for _, object := range l.objects {
		if object.hit(r, tMin, closest, &tempRec) {
			hitAnything = true
			closest = tempRec.t
			*rec = tempRec
		}
	}

	return hitAnything
}

func (l hittableList) boundingBox(time0 float64, time1 float64, box *aabb) bool {
	if len(l.objects) == 0 {
		return false
	}

	tempBox := aabb{}
	firstBox := true

	for _, object := range l.objects {
		if !object.boundingBox(time0, time1, &tempBox) {
			return false
		}
		if firstBox {
			*box = tempBox
		} else {
			surroundingBox(*box, tempBox)
		}
		firstBox = false
	}

	return true
}

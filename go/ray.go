package main

type ray struct {
	orig point3
	dir  vec3
	tm   float64
}

func newRay(origin point3, direction vec3, time float64) ray {
	return ray{orig: origin, dir: direction, tm: time}
}

func (r ray) origin() point3 {
	return r.orig
}

func (r ray) direction() vec3 {
	return r.dir
}

func (r ray) time() float64 {
	return r.tm
}

func (r ray) at(t float64) point3 {
	return add(r.orig, mult(t, r.dir))
}

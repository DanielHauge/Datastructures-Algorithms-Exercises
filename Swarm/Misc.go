package main

type Point struct {
	x float64
	y float64
}

type FPoint struct {
	x float64
	y float64
	u float64
	w float64
}

func EvaluatePoint(p Point) float64{
	inter := -p.x * p.x - p.y * p.y
	result := p.x * inter
	return result
}

func EvaluateFPoint(p FPoint) float64{
	return 2 * p.x * (-p.x*p.x - p.y * p.y - (p.u-1)*(p.u-1) - p.w * p.w)
}

package main

type fhive struct {
	GlobalBest FPoint
	GlobalPoint float64
	Phohive float64 // around 0.8
}



func (h *fhive) GetBroadcast(p FPoint){
	if EvaluateFPoint(p)>h.GlobalPoint{
		h.GlobalBest = p
		h.GlobalPoint = EvaluateFPoint(p)
	}
}

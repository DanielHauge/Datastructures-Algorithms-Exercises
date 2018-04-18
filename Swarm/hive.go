package main


type hive struct {
	GlobalBest Point
	GlobalPoint float64
	Phohive float64 // around 0.8
}



func (h *hive) GetBroadcast(p Point){
	if EvaluatePoint(p)>h.GlobalPoint{
		h.GlobalBest = p
		h.GlobalPoint = EvaluatePoint(p)
	}
}



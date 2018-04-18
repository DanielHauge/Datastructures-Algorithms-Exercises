package main

import (

	"time"
	"math/rand"
)

type fparticle struct {
	Ownbest FPoint
	Position FPoint
	inertia float64
	vmax float64
	phopat float64 // around 0.1
	curVx float64
	curVy float64
	curVu float64
	curVw float64
}

func (p *fparticle) move(){
	//log.Println(p.Position.x, p.Position.y)
	//log.Println("TO: ")
	p.Position = p.CalculateNewPoint()
	//log.Println(p.Position.x, p.Position.y)
	if EvaluateFPoint(p.Position)>EvaluateFPoint(p.Ownbest){
		p.Ownbest = p.Position
		fHive.GetBroadcast(p.Position)
		//log.Println("I found new ownbest")
	}

}

func (p *fparticle) CalculateNewPoint()FPoint{
	p.curVx = (p.inertia * p.curVx) + (p.phopat * (p.Ownbest.x - p.Position.x)) + (fHive.Phohive * (fHive.GlobalBest.x - p.Position.x))
	p.curVy = (p.inertia * p.curVy) + (p.phopat * (p.Ownbest.y - p.Position.y)) + (fHive.Phohive * (fHive.GlobalBest.y - p.Position.y))
	p.curVu = (p.inertia * p.curVu) + (p.phopat * (p.Ownbest.u - p.Position.u)) + (fHive.Phohive * (fHive.GlobalBest.u - p.Position.u))
	p.curVw = (p.inertia * p.curVw) + (p.phopat * (p.Ownbest.w - p.Position.w)) + (fHive.Phohive * (fHive.GlobalBest.w - p.Position.w))

	if p.curVy > p.vmax{
		p.curVy = p.vmax
	}
	if p.curVy > p.vmax{
		p.curVx = p.vmax
	}

	return FPoint{p.Position.x+p.curVx, p.Position.y+p.curVy, p.Position.u+p.curVu, p.Position.w+p.curVw}
}

func CreateFParticleAndGo(){
	point := FPoint{(rand.Float64()*4)-2, (rand.Float64()*4)-2, (rand.Float64()*4)-2, (rand.Float64()*4)-2}

	bee := fparticle{point, point, 0.1, 0.01, 0.1, 0,0, 0, 0}
	fHive.GetBroadcast(point)
	for {
		if fHive.GlobalBest != bee.Position{
			bee.move()
		}
		time.Sleep(time.Millisecond*100)
	}
}

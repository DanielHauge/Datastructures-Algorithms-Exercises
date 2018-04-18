package main

import (
	"math/rand"
	"time"
	"log"
)

type particle struct {
	Ownbest Point
	Position Point
	inertia float64
	vmax float64
	phopat float64 // around 0.1
	curVx float64
	curVy float64
}

func (p *particle) move(){
	//log.Println(p.Position.x, p.Position.y)
	//log.Println("TO: ")
	p.Position = p.CalculateNewPoint()
	//log.Println(p.Position.x, p.Position.y)
	if EvaluatePoint(p.Position)>EvaluatePoint(p.Ownbest){
		p.Ownbest = p.Position
		MasterHive.GetBroadcast(p.Position)
		//log.Println("I found new ownbest")
	}

}

func (p *particle) CalculateNewPoint()Point{
	p.curVx = (p.inertia * p.curVx) + (p.phopat * (p.Ownbest.x - p.Position.x)) + (MasterHive.Phohive * (MasterHive.GlobalBest.x - p.Position.x))
	p.curVy = (p.inertia * p.curVy) + (p.phopat * (p.Ownbest.y - p.Position.y)) + (MasterHive.Phohive * (MasterHive.GlobalBest.y - p.Position.y))

	if p.curVy > p.vmax{
		p.curVy = p.vmax
	}
	if p.curVy > p.vmax{
		p.curVx = p.vmax
	}

	return Point{p.Position.x+p.curVx, p.Position.y+p.curVy}
}

func CreateParticleAndGo(){
	point := Point{(rand.Float64()*4)-2, (rand.Float64()*4)-2}
	if (point.x < -2.0 || point.y < -2.0){
		log.Println("WHAT?")
	}
	bee := particle{point, point, 0.1, 0.01, 0.1, 0,0}
	MasterHive.GetBroadcast(point)
	for {
		if MasterHive.GlobalBest != bee.Position{
			bee.move()
		}
		time.Sleep(time.Millisecond*100)
	}
}
package main

import (
	"log"
	"time"
)

var MasterHive *hive
var fHive *fhive

func main() {
	/*
	point := Point{0,0}
	MasterHive = &hive{point, -100, 0.8}

	log.Println(EvaluatePoint(point))



	go CreateParticleAndGo()
	go CreateParticleAndGo()
	go CreateParticleAndGo()
	go CreateParticleAndGo()
	go CreateParticleAndGo()
	go CreateParticleAndGo()
	go CreateParticleAndGo()
	*/


	point := FPoint{0,0,0,0}
	fHive = &fhive{point, -100, 0.8}

	log.Println(EvaluateFPoint(point))



	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()
	go CreateFParticleAndGo()



	for {

		log.Println("Hello")

		/*
		log.Println(MasterHive.GlobalPoint)
		log.Println(MasterHive.GlobalBest.y, MasterHive.GlobalBest.x)
		*/
		log.Println(fHive.GlobalPoint)
		log.Println(fHive.GlobalBest.y, fHive.GlobalBest.x, fHive.GlobalBest.u, fHive.GlobalBest.w)

		time.Sleep(time.Second)
	}


}




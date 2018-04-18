package main

import (
	"log"
	"strconv"
	"time"
	"math/rand"
	"fmt"
	"sort"
	)


// Algorithm that learns to play NIM. Pick 1 2 or 3 points, the player who takes the last point wins.
// Change NIM count to alter it. Output in the end will say all the bad positions which the algorithm will learn to avoid in the future.

var nim int = 20
var start int = 0

func main() {
	rand.Seed(time.Now().Unix())
	p1 := player{map[int]bool{},0}
	p2 := player{map[int]bool{},0}

	for x:= 0; x<70; x++ {

		x := start
		fmt.Print("\n\nStarting up!:\n\n")

		for x < nim {
			x = p1.takeoptimalturn(x)
			fmt.Print(" (P1)-> " + strconv.Itoa(x))
			if x == nim {
				p2.badpositions[p2.lastmove] = true
				fmt.Print(" Player 1 wins!: Player 2 will remember this!!!: remembers position: " + strconv.Itoa(p2.lastmove))
				break;
			}
			x = p2.takeoptimalturn(x)
			fmt.Print(" (P2)-> " + strconv.Itoa(x))
			if x == nim {
				p1.badpositions[p1.lastmove] = true
				fmt.Print(" Player 2 wins!: Player 1 will remember this!!!: remembers position: " + strconv.Itoa(p1.lastmove))
			}
		}

	}
	fmt.Print("\n\nPlayer 1 bad positions\n")
	array1 := []int{}
	for x, _ := range p1.badpositions{
		array1 = append(array1, x)
	}

	sort.Ints(array1)

	for _, x := range array1{
		fmt.Print(x)
		fmt.Print(" ")
	}

	fmt.Print("\n\nPlayer 2 bad positions\n")

	array2 := []int{}
	for x, _ := range p2.badpositions{
		array2 = append(array2, x)
	}

	sort.Ints(array2)

	for _, x := range array2{
		fmt.Print(x)
		fmt.Print(" ")
	}



}



type player struct {
	badpositions map[int]bool
	lastmove int
}

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func (p *player) takeoptimalturn(x int)int{


	ran := random(1, 4)

	if nim==x+3{
		p.lastmove = x+3
		return x+3
	} else if nim==x+2{
		p.lastmove = x+2
		return x+2
	} else if nim==x+1{
		p.lastmove = x+1
		return x+1
	} else if ran == 3{
		if contains(p.badpositions, ran+x){
			if contains(p.badpositions, ran+x-1){
				if contains(p.badpositions, ran+x-2){
					fmt.Println(" | Urg, cannot avoid bad position now: "+strconv.Itoa(x)+" To "+strconv.Itoa(x+ran-2))
					return x+ran-2
				} else {
					p.lastmove = x+ran-2
					return x+ran-2
				}
			} else {
				p.lastmove = x+ran-1
				return x+ran-1
			}
		} else {
			p.lastmove = x+ran
			return x+ran
		}
	} else if ran == 2{
		if contains(p.badpositions, ran+x){
			if contains(p.badpositions, ran+x-1){
				if contains(p.badpositions, ran+x+1){
					fmt.Println(" | Urg, cannot avoid bad position now: "+strconv.Itoa(x)+" To "+strconv.Itoa(x+ran-1))
					return x+ran-1

				} else {
					p.lastmove = x+ran+1
					return x+ran+1
				}

			} else {
				p.lastmove = x+ran+1
				return x+ran-1
			}
		} else {
			p.lastmove = x+ran
			return x+ran
		}
	} else if ran == 1{
		if contains(p.badpositions, ran+x){
			if contains(p.badpositions, ran+x+1){
				if contains(p.badpositions, ran+x+2){
					fmt.Println(" | Urg, cannot avoid bad position now: "+strconv.Itoa(x)+" To "+strconv.Itoa(x+ran))
					return x+ran

				} else {
					p.lastmove = x+ran+2
					return x+ran+2
				}
			} else {
				p.lastmove = x+ran+1
				return x+ran+1
			}
		} else {
			p.lastmove = x+ran
			return x+ran
		}
	} else {
		log.Println("Uuuh something went wrong")
		return x+ran
	}


}


func contains(positons map[int]bool, searchInt int) bool {
	return positons[searchInt]
}

package main

import "log"

var choises int

func main() {
	choises = 0
	jump(0)
	log.Println(choises)
}

func jump(x int){

	if x == 10{
		jump(x+1)
	} else if x < 10 {
		jump(x+1)
		jump(x+2)
	} else if x == 11{
		choises+=1
	}
}

// 11 steps. jumps either 1 or 2 for every jump. How many ways?

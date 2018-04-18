package main

import "log"

var choises int

func main() {
	choises = 0
	jump(0)
	log.Println(choises)
}

func jump(x int){

	if x == 17{
		jump(x+1)
		jump(x+2)
		jump(x+3)
	} else if x== 18{
		jump(x+1)
		jump(x+2)
	} else if x== 19{
		jump(x+1)
	} else if x <17 {
		jump(x+1)
		jump(x+2)
		jump(x+3)
		jump(x+4)
	} else if x == 20{
		choises+=1
	}
}

// 20 steps on the stair. jumps either 1, 2, 3 or 4 for every jump. How many ways can you get up the stairs.

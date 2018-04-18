package main

import "log"

func main() {
	log.Println(tilecal(10))
}

func tilecal(x int)int{
	if x==0 {return 1}
	if x == 1 {return 1}
	if x == 2 {return 2}
	return tilecal(x-1) + tilecal(x-2) // 2*tilecal(x-1)
}

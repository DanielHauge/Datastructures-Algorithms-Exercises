package main

import "log"

func main() {
	log.Println(Treeways(3))
}

func Treeways(x int)int{
	if x == 0 { return 1}
	if x == 1 { return 1}
	if x == 2 {return 2}
	return Treeways(x-2)*Treeways(x-1)+Treeways(x-1)*Treeways(x-2)
}



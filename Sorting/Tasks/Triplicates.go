package main

import (
	"log"
)

type maps struct {
	maps []map[string]bool
}

var(
	l1 = []string {"Daniel", "Kristian", "Emmely", "Yoyo", "whaatup", "Eeharg?"}
	l2 = []string {"Tobias", "Keven", "Emily", "Yo", "whap", "Daniel"}
	l3 = []string {"Daniel", "Kristian", "Emmely", "Yoyo", "whaatup", "Eeharg?"}
	m = maps{}
)

func main() {
	log.Println(FirstNameInCommon(l1,l2,l3))
}

func FirstNameInCommon(arrays ...[]string)string{


	m.maps = append(m.maps, make(map[string]bool))
	m.maps = append(m.maps, make(map[string]bool))
	m.maps = append(m.maps, make(map[string]bool))
	temp := "none"
	for i := 0; (i<len(arrays[0])||i<len(arrays[1])||i<len(arrays[2])); i++ {
		temp = SetTrue(m.maps[0], arrays[0][i])
		if temp != "none"{
			return temp
		}
		temp = SetTrue(m.maps[1], arrays[1][i])
		if temp != "none"{
			return temp
		}
		temp = SetTrue(m.maps[2], arrays[2][i])
		if temp != "none"{
			return temp
		}
	}

	return "none"
}

func SetTrue(ma map[string]bool, name string)string{


	ma[name] = true
	r := true
	for _, v := range m.maps {
		if !v[name]{
			r = false
		}
	}
	if r{
		return name
	}else {
		return "none"
	}

}

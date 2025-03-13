package main

import "fmt"

func UpdateSlice(sl []string, s string) {
	if len(sl) == 0 {
		return
	}
	sl[len(sl)-1] = s
	fmt.Println(sl)
}

func GrowSlice(sl []string, s string) {
	sl = append(sl, s)
	fmt.Println(sl)
}

func main() {
	sl := []string{"apple", "banana", "orange"}
	fmt.Println(sl)
	UpdateSlice(sl, "mango")
	fmt.Println(sl)
	GrowSlice(sl, "pomegranate")
	fmt.Println(sl)
}

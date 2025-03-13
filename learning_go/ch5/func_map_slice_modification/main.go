package main

import "fmt"

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}
func modSlice(s []int) {
	// Modifies original slice
	fmt.Printf("val:%v, address:%p, len: %v, cap: %v\n", s, &s, len(s), cap(s))
	for k, v := range s {
		s[k] = v * 2
	}
	fmt.Printf("val:%v, address:%p, len: %v, cap: %v\n", s, &s, len(s), cap(s))
	s = append(s, 10)
	fmt.Printf("val:%v, address:%p, len: %v, cap: %v\n", s, &s, len(s), cap(s))
	for k, v := range s {
		s[k] = v * 2
	}
	fmt.Printf("val:%v, address:%p, len: %v, cap: %v\n", s, &s, len(s), cap(s))
}

func main() {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)
	//s := []int{1, 2, 3} // when append it will allocate newArray in modSlice
	s := make([]int, 3, 4)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Printf("val:%v, address:%p, len: %v, cap: %v\n", s, &s, len(s), cap(s))
	modSlice(s)
	fmt.Printf("val:%v, address:%p, len: %v, cap: %v\n", s, &s, len(s), cap(s))
}

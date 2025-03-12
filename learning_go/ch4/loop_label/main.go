package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_π!", "mate", "sick", "paper"}
	// Use labels if need to break and continue in nested loops
	// Its basically goto(unconditional jump)
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
			if r == 'i' {
				break outer
			}
		}
		fmt.Println()
	}

	// Break from loop inside switch using labels
	for i := range 10 {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break
		default:
			fmt.Println(i, "is boring")
		}
	}
loop:
	for i := range 10 {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}

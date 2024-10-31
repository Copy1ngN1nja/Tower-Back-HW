package main

import "fmt"

func make0(x *int64, bit int64) {
	*x &= ^(1 << bit)
}

func make1(x *int64, bit int64) {
	*x |= (1 << bit)
}

func main() {
	fmt.Println("Enter your number")
	var x int64
	fmt.Scan(&x)
	for {
		fmt.Println("If you want to set the ith bit to a 0, type 'i 0', where i is a number of bit")
		fmt.Println("If you want to set the ith bit to a 1, type 'i 1', where i is a number of bit")
		fmt.Println("If you want to exit program, type -1")
		var bit, val int64
		fmt.Scan(&bit)
		if bit == -1 {
			break
		}
		fmt.Scan(&val)
		if val == 0 {
			make0(&x, bit)
		} else {
			make1(&x, bit)
		}
		fmt.Printf("x = %d\n", x)
	}
}

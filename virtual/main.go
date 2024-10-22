package main

import "fmt"

func main() {
	move(4, "A", "B", "C")
}

func move(n int, src, dest, buffer string) {
	if n == 1 {
		fmt.Printf("%d : %s  --> %s\n", n, src, dest)
		return
	}
	move(n-1, src, buffer, dest)
	fmt.Printf("%d : %s  --> %s\n", n, src, dest)
	move(n-1, buffer, dest, src)
}

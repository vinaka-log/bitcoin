package main

import"fmt"

func init() {
	fmt.Println("init!")
}

func bazz() {
	fmt.Println("BAZZ")
}

func main() {
	bazz()
	fmt.Println("Hello World", "TEST TEST")
}

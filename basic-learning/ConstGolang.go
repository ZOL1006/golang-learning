package main

import "fmt"
const (
	i=1<<iota
	j
	k=4<<iota
	l
	m=iota
	n
)

func main() {
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
	fmt.Println("m=",m)
	fmt.Println("n=",n)
}
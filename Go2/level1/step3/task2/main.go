package main

func Receive(ch chan int) int {
	var num int
	num = <-ch
	return num
}

func main() {
}

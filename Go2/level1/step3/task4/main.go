package main

func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for _, i := range nums {
		ch <- i
	}
	return ch
}
func main() {
}

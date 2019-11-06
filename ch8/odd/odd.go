package odd

func PrintOdd(n int) []int {
	var res []int
	ch := make(chan int, 1)
	for i := 0; i < n; i++ {
		select {
		case ch <- i:
		case num := <-ch:
			res = append(res, num)
		}
	}
	return res
}

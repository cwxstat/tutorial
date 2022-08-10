package simple

func NewChanOwner(num int) func() <-chan int {
	ch := make(chan int, num)
	chanOwner := func() <-chan int {

		go func() {
			defer close(ch)
			for i := 0; i < num; i++ {
				ch <- i
			}
		}()
		return ch
	}
	return chanOwner
}

func Consumer(chanOwner func(int) func() <-chan int, num int) {
	c := chanOwner(num)
	resultStream := c()
	for i := range resultStream {
		println(i)
	}
}

package simple

func NewChanOwner(num int) func() <-chan int {
	ch := make(chan int)
	chanOwner := func() <-chan int {
		go func() {
			for i := 0; i < num; i++ {
				ch <- i
			}
			close(ch)
		}()
		return ch
	}
	return chanOwner
}

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

// NewChanOwner returns a function that returns a channel.
func NewChanOwnerFunc(f func(interface{}) interface{}, arg interface{}, num int) func() <-chan interface{} {
	ch := make(chan interface{}, num)
	chanOwner := func() <-chan interface{} {

		go func() {
			defer close(ch)
			for i := 0; i < num; i++ {
				ch <- f(arg)
			}
		}()
		return ch
	}
	return chanOwner
}

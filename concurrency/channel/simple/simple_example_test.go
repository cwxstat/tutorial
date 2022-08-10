package simple_test

import (
	"github.com/cwxstat/tutorial/concurrency/channel/simple"
)

func ExampleConsumer() {
	simple.Consumer(simple.NewChanOwner, 3)

	// Output:

}

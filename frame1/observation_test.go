package frame1

import "fmt"

func ExampleNewObservation() {
	o := newObservation()
	fmt.Printf("test: NewObservation() -> [lat:%v] [grad:%v]\n", o.Latency, o.Gradient)

	o = newObservation()
	fmt.Printf("test: NewObservation() -> [lat:%v] [grad:%v]\n", o.Latency, o.Gradient)

	//Output:
	//fail
}

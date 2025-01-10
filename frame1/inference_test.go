package frame1

import "fmt"

func ExampleInference() {
	f := Inference(Observation{Latency: 100, Gradient: 5})
	fmt.Printf("test: Inference() -> [sat:%v] [grad:%v]\n", f.Impress.Saturation, f.Impress.Gradient)

	f = Inference(Observation{Latency: 1100, Gradient: 5})
	fmt.Printf("test: Inference() -> [sat:%v] [grad:%v]\n", f.Impress.Saturation, f.Impress.Gradient)

	f = Inference(Observation{Latency: 1900, Gradient: 5})
	fmt.Printf("test: Inference() -> [sat:%v] [grad:%v]\n", f.Impress.Saturation, f.Impress.Gradient)

	f = Inference(Observation{Latency: 100, Gradient: 5})
	fmt.Printf("test: Inference() -> [sat:%v] [grad:%v]\n", f.Impress.Saturation, f.Impress.Gradient)

	f = Inference(Observation{Latency: 100, Gradient: 9})
	fmt.Printf("test: Inference() -> [sat:%v] [grad:%v]\n", f.Impress.Saturation, f.Impress.Gradient)

	f = Inference(Observation{Latency: 100, Gradient: 15})
	fmt.Printf("test: Inference() -> [sat:%v] [grad:%v]\n", f.Impress.Saturation, f.Impress.Gradient)

	//Output:
	//test: Inference() -> [sat:low] [grad:low]
	//test: Inference() -> [sat:medium] [grad:low]
	//test: Inference() -> [sat:high] [grad:low]
	//test: Inference() -> [sat:low] [grad:low]
	//test: Inference() -> [sat:low] [grad:medium]
	//test: Inference() -> [sat:low] [grad:high]

}

package frame1

import "fmt"

func ExampleFrames() {
	f1 := NewFrame(newObservation())
	comprehension(f1)
	addFrame(f1)

	f2 := NewFrame(newObservation())
	comprehension(f2)
	addFrame(f2)

	fmt.Printf("test: Frames() -> [sat:%v] [grad:%v] [stars:%v]\n", f1.Impress.Saturation, f1.Impress.Gradient, f1.Stars)
	fmt.Printf("test: Frames() -> [sat:%v] [grad:%v] [stars:%v]\n", f2.Impress.Saturation, f2.Impress.Gradient, f2.Stars)

	updateFeedback()
	fmt.Printf("test: Frames() -> [sat:%v] [grad:%v] [stars:%v]\n", f1.Impress.Saturation, f1.Impress.Gradient, f1.Stars)
	fmt.Printf("test: Frames() -> [sat:%v] [grad:%v] [stars:%v]\n", f2.Impress.Saturation, f2.Impress.Gradient, f2.Stars)

	//Output:
	//fail

}

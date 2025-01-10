package frame1

import (
	"context"
	"fmt"
	"github.com/behavioral-ai/core/core"
)

const (
	PkgPath = "github/behavioral-ai/ingress/frame1"
)

func Observe(ctx context.Context, origin core.Origin) (Observation, *core.Status) {
	return newObservation(), core.StatusOK()
}

func AddAction(ctx context.Context, origin core.Origin, action int) *core.Status {
	fmt.Printf("Action: %v %v\n", origin, action)
	return core.StatusOK()
}

func AddInference(ctx context.Context, origin core.Origin, f *Frame) *core.Status {
	if f != nil {
		addFrame(f)
		fmt.Printf("Frame: %v %v\n", origin, f.Observe)
	}
	return core.StatusOK()
}

func Inference(o Observation) *Frame {
	f := NewFrame(o)
	comprehension(f)
	reasoning(f)
	return f
}

func AddFeedback(ctx context.Context) {
	updateFeedback()
}

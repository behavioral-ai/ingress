package frame1

import (
	"math/rand"
	"sync"
)

const (
	low    = "low"
	medium = "medium"
	high   = "high"
)

type Impression struct {
	Saturation string
	Gradient   string
}

type Frame struct {
	Observe     Observation
	Impress     Impression
	ReasonId    string
	RateLimit   int // Percentage of traffic
	Stars       int
	StarsReason string
}

func NewFrame(o Observation) *Frame {
	f := new(Frame)
	f.Observe = o
	return f
}

var (
	frames []*Frame
	muf    = sync.Mutex{}
)

func addFrame(f *Frame) {
	muf.Lock()
	defer muf.Unlock()
	frames = append(frames, f)
}

func findFrame(i Impression) *Frame {
	muf.Lock()
	defer muf.Unlock()
	var f *Frame
	for _, frame := range frames {
		if frame.Impress.Gradient != i.Gradient || frame.Impress.Saturation != i.Saturation {
			continue
		}
		if f == nil {
			f = frame
			continue
		}
		if frame.Stars > f.Stars {
			f = frame
		}
	}
	return f
}

func updateFeedback() {
	muf.Lock()
	minN := 1
	maxN := 5
	for _, frame := range frames {
		if frame.Stars == 0 {
			frame.Stars = rand.Intn(maxN-minN+1) + minN
		}
	}
}

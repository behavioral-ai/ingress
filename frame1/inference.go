package frame1

import "sync"

// observation -> comprehend -> impression
// impression -> reason -> action

const (
	threshold = 2500
)

type lookup struct {
	low    int
	medium int
	high   int
}

var (
	sat  = lookup{low: 25, medium: 65, high: 80}
	grad = lookup{low: 5, medium: 10, high: 15}
)

func comprehension(f *Frame) {
	if f == nil {
		return
	}
	f.Impress.Saturation = lookupSaturation(f.Observe.Latency)
	f.Impress.Gradient = lookupGradient(f.Observe.Gradient)
}

func lookupSaturation(latency int) string {
	s := float64(latency) / float64(threshold)
	s *= 100
	if s <= float64(sat.low) {
		return low
	}
	if s <= float64(sat.medium) {
		return medium
	}
	return high
}

func lookupGradient(gradient int) string {
	if gradient <= grad.low {
		return low
	}
	if gradient <= grad.medium {
		return medium
	}
	return high
}

type reason struct {
	id         string
	saturation string
	gradient   string
	action     int
}

var (
	mur = sync.Mutex{}

	reasons = []reason{
		{id: "low-low-1", saturation: low, gradient: low, action: 0},
		{id: "low-medium-1", saturation: low, gradient: medium, action: 0},
		{id: "low-high-1", saturation: low, gradient: high, action: 10},

		{id: "medium-low-1", saturation: medium, gradient: low, action: 0},
		{id: "medium-medium-1", saturation: medium, gradient: medium, action: 10},
		{id: "medium-high-1", saturation: medium, gradient: high, action: 20},

		{id: "high-low-1", saturation: high, gradient: low, action: 10},
		{id: "high-medium-1", saturation: high, gradient: medium, action: 20},
		{id: "high-high-1", saturation: high, gradient: high, action: 40},
		/*
			{id: "high-low-2", saturation: high, gradient: low, action: 15},
			{id: "high-medium-2", saturation: high, gradient: medium, action: 30},
			{id: "high-high-2", saturation: high, gradient: high, action: 60},

		*/
	}
)

func findReason(i Impression) reason {
	mur.Lock()
	defer mur.Unlock()
	for _, r := range reasons {
		if r.gradient == i.Gradient && r.saturation == i.Saturation {
			return r
		}
	}
	return reason{}
}

func selectReason(id string) reason {
	mur.Lock()
	defer mur.Unlock()
	for _, r := range reasons {
		if r.id == id {
			return r
		}
	}
	return reason{}
}

func reasoning(f *Frame) {
	if f == nil {
		return
	}
	//r := inferReason(f)
	//f.Action = r.action
}

func inferReason(f *Frame) reason {
	// First check inference history to see if this impression has been reasoned on before
	frame := findFrame(f.Impress)
	// If the impression has not been reasoned on, then select the reason from the defaults
	if frame == nil {
		return findReason(f.Impress)
	}
	// If there are no stars, then except
	if frame.Stars == 0 {
		return selectReason(frame.ReasonId)
	}
	// Now determine if this reason is good enough, or if there is a way to create a new better reason
	return reason{}
}

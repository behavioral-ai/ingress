package operative1

import (
	"fmt"
	"github.com/behavioral-ai/core/core"
)

func ExampleNewAgent() {
	a := NewAgent(core.Origin{Region: "us-central"}, nil, nil)

	fmt.Printf("test: NewAgent() -> [uri:%v]\n", a)

	//Output:
	//test: NewAgent() -> [uri:operative1:us-central..]

}

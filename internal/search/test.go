package search

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func main() {
	jsFile := "./secret.js"
	bytes, err := ioutil.ReadFile(jsFile)
	if err != nil {
		panic(err)
	}

	vm := otto.New()
	_, _ = vm.Run(string(bytes))
	enc, err := vm.Call("b", "12345")
	fmt.Println(enc)
}

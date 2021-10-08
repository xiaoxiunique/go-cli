/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import "go-cli/cmd"

func main() {
	cmd.Execute()
}

//func main() {
//	jsFile := "./secret.js"
//	bytes, err := ioutil.ReadFile(jsFile)
//	if err != nil {
//		panic(err)
//	}
//
//	vm := otto.New()
//	_, err = vm.Run(string(bytes))
//	if err != nil {
//		panic(err)
//	}
//	enc, err := vm.Call("b", nil, "123456")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(enc.String())
//}

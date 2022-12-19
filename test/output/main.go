package main

import (
	json "encoding/json"
	os "os"

	hello "github.com/mikerybka/execserver/test/pkg/hello"
)

func main() {
	var input struct {
		S string `json:"s"`
	}
	json.NewDecoder(os.Stdin).Decode(&input)
	out1, out2 := hello.Echo(input.S)
	var output struct {
		Out1 string `json:"out1"`
		Out2 error  `json:"out2"`
	}
	output.Out1 = out1
	output.Out2 = out2
	json.NewEncoder(os.Stdout).Encode(output)
}

package main

import (
	hello "github.com/mikerybka/execserver/test/pkg/hello"
	json "encoding/json"
	os "os"
)

func main() {
	var input struct {
	}
	json.NewDecoder(os.Stdin).Decode(&input)
	out1 := hello.Hello()
	var output struct {
		Out1 string `json:"out1"`
	}
	output.Out1 = out1
	json.NewEncoder(os.Stdout).Encode(output)
}

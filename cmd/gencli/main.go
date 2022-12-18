package main

import (
	"fmt"

	"github.com/mikerybka/execserver/pkg/golang"
)

func main() {
	srcDir := "../../../"
	pkg := "github.com/mikerybka/execserver/test/pkg/hello"
	fn := "Hello"
	mainFile := "test/output/main.go"
	err := golang.GenerateCLI(srcDir, pkg, fn, mainFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

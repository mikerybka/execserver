package golang

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/library-development/go-nameconv"
)

// GenerateCLI geneates a simple Go CLI app that only calls a single func.
// The resulting command has no arguments.
// It reads function input from stdin in JSON format and writes the result to stdout in JSON format.
func GenerateCLI(srcDir, pkg, funcName, outFile string) error {
	funcSignature, err := ReadFuncSignature(srcDir, pkg, funcName)
	if err != nil {
		return err
	}
	importMap := ImportMap{}
	importMap.AddPackage(pkg)
	importMap.AddPackage("encoding/json")
	importMap.AddPackage("os")
	for _, input := range funcSignature.Inputs {
		importMap.AddPackage(input.Type.Pkg)
	}
	for _, output := range funcSignature.Outputs {
		importMap.AddPackage(output.Type.Pkg)
	}
	err = os.MkdirAll(filepath.Dir(outFile), os.ModePerm)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	b.WriteString("package main\n\n")
	importMap.Write(&b)
	b.WriteString("\nfunc main() {\n")
	b.WriteString("\tvar input struct {\n")
	for _, input := range funcSignature.Inputs {
		n := nameconv.ParsePascalCase(input.Name)
		b.WriteString("\t\t")
		b.WriteString(n.PascalCase())
		b.WriteString(" ")
		b.WriteString(input.Type.Name)
		b.WriteString(" `json:\"")
		b.WriteString(n.SnakeCase())
		b.WriteString("\"\n")
	}
	b.WriteString("\t}\n")
	b.WriteString("}\n")
	b.WriteString("\t\n")
	return os.WriteFile(outFile, b.Bytes(), os.ModePerm)
}

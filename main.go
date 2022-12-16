package main

import (
	"flag"
	"net/http"

	"github.com/mikerybka/execserver/pkg/web"
)

func main() {
	flag.Parse()
	port := flag.Arg(0)
	authDir := flag.Arg(1)
	srcDir := flag.Arg(2)
	s := web.ExecServer{
		AuthDir:   authDir,
		SourceDir: srcDir,
	}
	http.Handle("/", &s)
	http.ListenAndServe(":"+port, nil)
}

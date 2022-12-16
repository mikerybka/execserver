package main

import (
	"flag"
	"net/http"
)

func main() {
	flag.Parse()
	port := flag.Arg(0)
	authDir := flag.Arg(1)
	srcDir := flag.Arg(2)
	s := Server{
		AuthDir:   authDir,
		SourceDir: srcDir,
	}
	http.Handle("/", &s)
	http.ListenAndServe(":"+port, nil)
}

type Server struct {
	AuthDir   string
	SourceDir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

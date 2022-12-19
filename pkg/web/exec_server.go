package web

import (
	"encoding/json"
	"io"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mikerybka/execserver/pkg/golang"
)

type ExecServer struct {
	AuthDir   string
	SourceDir string
}

func (s *ExecServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	auth := &Auth{Dir: s.AuthDir}
	if !auth.IsAdmin(req.Token) {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}
	mainFile := filepath.Join(s.SourceDir, "cmd", req.Pkg, strings.ToLower(req.Func), "main.go")
	err := golang.GenerateCLI(s.SourceDir, req.Pkg, req.Func, mainFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cmd := exec.Command("go", "run", mainFile)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		http.Error(w, "stdin pipe", http.StatusInternalServerError)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		http.Error(w, "stdout pipe", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(stdin).Encode(req.Inputs)
	if err != nil {
		http.Error(w, "encode", http.StatusInternalServerError)
		return
	}
	err = stdin.Close()
	if err != nil {
		http.Error(w, "close", http.StatusInternalServerError)
		return
	}
	err = cmd.Start()
	_, err = io.Copy(w, stdout)
	if err != nil {
		http.Error(w, "copy", http.StatusInternalServerError)
		return
	}
}

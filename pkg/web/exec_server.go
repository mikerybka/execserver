package web

import (
	"encoding/json"
	"fmt"
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
		http.Error(w, fmt.Sprintf("stdin pipe: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		http.Error(w, fmt.Sprintf("stdout pipe: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(stdin).Encode(req.Inputs)
	if err != nil {
		http.Error(w, fmt.Sprintf("encode: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	err = stdin.Close()
	if err != nil {
		http.Error(w, fmt.Sprintf("stdin close: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	err = cmd.Start()
	if err != nil {
		http.Error(w, fmt.Sprintf("start: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	b, err := io.ReadAll(stdout)
	if err != nil {
		http.Error(w, fmt.Sprintf("read: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

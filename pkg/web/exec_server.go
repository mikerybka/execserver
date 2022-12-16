package web

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		http.Error(w, fmt.Errorf("not authorized").Error(), http.StatusUnauthorized)
		return
	}
	// TODO: generate main.go for the func
	// TODO: go run main.go
}

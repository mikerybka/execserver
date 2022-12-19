package web

import (
	"os"
	"path/filepath"
)

type Auth struct {
	Dir string
}

func (a *Auth) IsAdmin(token string) bool {
	sessionpath := filepath.Join(a.Dir, "sessions", token)
	adminpath := filepath.Join(a.Dir, "admin")
	_, err := os.Stat(sessionpath)
	if err != nil {
		return false
	}
	userID, err := os.ReadFile(sessionpath)
	if err != nil {
		return false
	}
	adminID, err := os.ReadFile(adminpath)
	if err != nil {
		return false
	}
	return string(userID) == string(adminID)
}

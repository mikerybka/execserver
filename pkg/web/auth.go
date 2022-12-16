package web

type Auth struct {
	Dir string
}

func (a *Auth) IsAdmin(token string) bool {
	return false
}

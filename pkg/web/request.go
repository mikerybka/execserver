package web

type Request struct {
	Token  string
	Pkg    string
	Func   string
	Inputs map[string]any
}

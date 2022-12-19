package hello

func Hello() string {
	return "Hello, world."
}

func Echo(s string) (string, error) {
	return s, nil
}

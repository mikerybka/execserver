package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Exec executes a Go function on an execserver.
func Exec[InputType, OutputType map[string]any](serverEndpoint, token, pkg, fn string, inputs InputType) (OutputType, error) {
	req := Request{
		Token:  token,
		Pkg:    pkg,
		Func:   fn,
		Inputs: inputs,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(serverEndpoint, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(resBody))
	}
	var output OutputType
	err = json.Unmarshal(resBody, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

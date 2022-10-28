package main

import (
	"context"
	"github.com/corezoid/gitcall-go-runner/gitcall"
)

func usercode(_ context.Context, data map[string]interface{}) error {

	data["hello"] = "Hello world 2"
	
	if data["key1"] == "val1 {
		data["test1"] = 123
	}

	return nil
}

func main() {
	gitcall.Handle(usercode)
}

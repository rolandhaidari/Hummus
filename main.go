package main

import (
	"github.com/Azer0s/Hummus/runner"
)

func main() {
	runner.RunRepl()
	runner.RunFile("examples/identity.hummus")
}

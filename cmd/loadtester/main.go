package main

import (
	"github.com/anagfranco/goexpert-stress-test/internal/infrastructure"
)

func main() {
	container := infrastructure.NewContainer()

	cli := container.GetCLI()

	cli.Execute()
}

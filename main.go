package main

import (
	"github.com/bagastri07/antarupa-test/application"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	figure.NewColorFigure("anantarupa", "", "blue", true).Print()

	app := application.NewApp()

	app.Start()
}

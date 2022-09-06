package main

import (

	// Dependencies of Turbine
	"github.com/meroxa/turbine-go"
	"github.com/meroxa/turbine-go/runner"
)

func main() {
	runner.Start(App{})
}

var _ turbine.App = (*App)(nil)

type App struct{}

func (a App) Run(v turbine.Turbine) error {
	cck, err := v.Resources("cck")
	if err != nil {
		return err
	}

	rr, err := cck.Records("topic_1", nil)
	if err != nil {
		return err
	}

	err = cck.Write(rr, "topic_2")
	if err != nil {
		return err
	}

	return nil
}

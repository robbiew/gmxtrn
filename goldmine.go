package main

import (
	"os"
)

type goTelnet struct{}

func newGoTelnet() *goTelnet {
	return new(goTelnet)
}

func (g *goTelnet) run() {
	telnetClient := g.createTelnetClient()
	telnetClient.ProcessData(os.Stdin, os.Stdout)
}

func (g *goTelnet) createTelnetClient() *client.TelnetClient {
	commandLine := commandline.Read()
	telnetClient := client.NewTelnetClient(commandLine)
	return telnetClient
}

package main

import (
	server "github.com/brainyard-io/kato/server"
)

// Kato type
type Kato struct {
	mode bool
}

// main runs the go program
func main() {
	kato := new(Kato)
	kato.Init()
	kato.Run()
}

// Init initiates kato
func (k *Kato) Init() {
	return
}

// Run runs kato
func (k *Kato) Run() {
	if k.mode {
		server.NewServer(server.ServerArgs{}).Serve()
	}
	return
}

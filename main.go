package main

import (
	server "github.com/brainyard-io/kato/server"
)

// the kato type
type Kato struct {
	mode	bool
}

// the main function
func main() {
	kato := new(Kato)
	kato.Init()
	kato.Run()
}

// Initiates kato
func (k *Kato) Init() {
	return
}

// Runs kato
func (k *Kato) Run() {
	if(k.mode) {
		server.NewServer(server.ServerArgs{}).Serve()
	}
	return
}
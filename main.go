package main

import (
	server "github.com/brainyard-io/kato/server"
)

type Kato struct {
	mode	bool
}

func main() {
	kato := new(Kato)
	kato.Init()
	kato.Run()
}

func (k *Kato) Init() {
	return
}

func (k *Kato) Run() {
	if(k.mode) {
		server.NewServer(server.ServerArgs{}).Serve()
	}
	return
}
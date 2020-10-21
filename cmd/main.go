package main

import "github.com/andrewGIS/apiserver/api"

// TODO Make simple view web

// Run server
func main() {
	s := api.Server{}
	s.Start()
}

package main

import "github.com/andrewGIS/apiserver/api"

// Run server
func main() {
	s := api.Server{}
	s.Start()
}

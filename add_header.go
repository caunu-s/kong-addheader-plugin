package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

func main() {
	server.StartServer(New, Version, Priority)
}

const Version = "1.0.0"
const Priority = 1

type MyConfig struct {
	FieldValue string
 }

func New() interface{} {
	return &MyConfig{}
}

func (conf MyConfig) Access(kong *pdk.PDK) {
	value := conf.FieldValue
	if value == "" {
		value = "bar"
	}
	kong.Response.AddHeader("foo", value)
}

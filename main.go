package main

import (
	"log"
	"plugin"
)

func main() {
	_, err := plugin.Open("plugin/hello.so")
	if err != nil {
		log.Fatal(err)
	}
	_, err = plugin.Open("plugin_v1/bye.so")
	if err != nil {
		log.Fatal(err)
	}

}

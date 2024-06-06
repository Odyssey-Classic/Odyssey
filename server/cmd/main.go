package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
)

func main() {
	var registry string
	flag.StringVar(&registry, "registry", "http://local.fosteredgames.com:8080", "Registry URL")
	flag.Parse()

	host, err := url.Parse(registry)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(host)
}

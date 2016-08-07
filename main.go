package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"os"
)

func main() {
	mc := memcache.New(os.Args...)
	if err := mc.Set(&memcache.Item{Key: "__mcping", Value: []byte("1")}); err != nil {
		log.Fatal(err)
	}
	// we need no more
}

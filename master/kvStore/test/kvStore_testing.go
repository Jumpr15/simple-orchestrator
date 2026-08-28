package main

import(
	"simple-orchestrator/master/kvStore"
	"log"
	"fmt"
)

func main() {
	kv := kvStore.New()

	kv.Set("Key", "Value")
	c, b := kv.Get("Key")
	d, a := kv.Get("Value")

	fmt.Println(c, d, b, a)

	kv.Del("Key")
	g, h := kv.Get("Key")
	fmt.Println(g, h)

	log.Printf("Hello kv")
}
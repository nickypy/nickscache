package main

import (
	"fmt"
)

func main() {
	app := App{NewCache(256)}
	server := BuildServer(&app)
	defer server.ListenAndServe()
	fmt.Printf("Listening on %s", server.Addr)
}

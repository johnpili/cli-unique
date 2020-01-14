package main

import "github.com/rs/xid"

import "fmt"

func main() {
	fmt.Println(generatePostID())
}

func generatePostID() string {
	guid := xid.New()
	guid.Time().Add(100)
	return guid.String()[12:20]
}

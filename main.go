package main

import (
	"github.com/dexguitar/chatapp/cmd"
	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()
}

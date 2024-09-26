package main

import (
	"fmt"
	"log/slog"

	"github.com/dexguitar/chatapp/cmd"
	_ "github.com/lib/pq"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		slog.Error(fmt.Sprintf("error executing app: %s", err.Error()))
		panic(err)
	}
}

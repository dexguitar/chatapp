package main

import (
	"fmt"

	"github.com/dexguitar/chatapp/app"
	"github.com/dexguitar/chatapp/configs"
	_ "github.com/lib/pq"
)

func main() {
	c, err := configs.LoadConfig(".env")
	if err != nil {
		panic(fmt.Sprintf("cannot load config: %s", err.Error()))
		return
	}

	app.Run(c)
}

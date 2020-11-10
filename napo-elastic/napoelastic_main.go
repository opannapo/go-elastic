package main

import (
	"flag"
	"fmt"
	"napoelastic/napoelastic"
)

const msg = `- Environment
	└ env			» %s	
`

func main() {
	env := flag.String("env", "local", "available --> local | dev | prod")
	flag.Parse()

	fmt.Printf(msg, *env)

	app := new(napoelastic.App)
	if *env == "dev" {
		app.ConfigFile = "config-dev.json"
	} else if *env == "prod" {
		app.ConfigFile = "config-prod.json"
	} else {
		app.ConfigFile = "config.json"
	}

	app.Run()
}

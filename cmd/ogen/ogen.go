package main

import (
	"github.com/kindynosmx/ogen/cmd/ogen/commands"
	_ "net/http/pprof"
)

func main() {
	err := commands.Execute()
	if err != nil {
		panic(err)
	}
}

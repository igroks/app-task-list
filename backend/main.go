package main

import (
	"flag"
	"fmt"

	"github.com/igroks/sd-project/backend/app/routes"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8380, "server listening port")
	flag.Parse()

	routes.HandleResquest().Run(fmt.Sprintf(":%d", port))
}

package main

import "flag"

func handleCLI() *string {
	portArgument := flag.String("port", "5101", "sets the port used by the web server")
	flag.Parse()

	return portArgument
}

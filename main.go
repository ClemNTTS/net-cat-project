package main

import (
	"os"
	"serv"
)

func main() {
	var ip, port string
	if len(os.Args) == 1 {
		ip = "localhost"
		port = "8080"
	} else if len(os.Args) == 2 {
		ip = "localhost"
		port = os.Args[1]
	} else {
		ip = os.Args[1]
		port = os.Args[2]
	}

	serv.StartServ(ip, port)
}

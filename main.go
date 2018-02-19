package main

import (
  "flag"
  "os"
  "./lib"
)

func main() {
  var isHost bool

  flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
  flag.Parse()

  if isHost {
  	// go run main.go -listen <ip>
  	connIp := os.Args[2] 
  	lib.RunHost(connIp)
  } else {
  	// go run main.go <ip>
  	connIp := os.Args[1]
  	lib.RunGuest(connIp)
  }
}


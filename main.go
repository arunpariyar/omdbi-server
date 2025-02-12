package main

import (
	"github.com/arunpariyar/omdbi-server/server"
	"github.com/arunpariyar/omdbi-server/utils"
)

func main() {
	config := utils.GetEnv()
	server := server.NewServer(config["apiKey"]); 
	server.StartServer()
}

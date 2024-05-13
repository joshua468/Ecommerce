package main

import (
	"log"

	"github.com/joshua468/bank-app/Desktop/Ecom/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

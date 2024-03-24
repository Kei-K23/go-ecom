package main

import "github.com/Kei-K23/go-ecom/cmd/api"

func main() {
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		panic(err)
	}
}

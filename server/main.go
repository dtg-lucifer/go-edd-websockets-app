package main

import (
	"fmt"

	"github.com/dtg-lucifer/go-eda-websocket-app/api"
)

func main() {
	api := api.NewApi()

	api.Init("/api/v1")
	if err := api.Start(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}

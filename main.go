package main

import (
	"fmt"
)

func main() {
	// Example usage

	app, err := initApp()

	if err != nil {
		fmt.Println("Error:", err)
	}

	app.Router.StaticFile("/", "./public/index.html")

	app.Run()
	defer app.Stop()

}

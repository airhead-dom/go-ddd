package main

import "go-ddd/di"

func main() {
	app := di.Init()
	app.Run()
}

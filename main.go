package main

import (
	"praktikum/config"
	m "praktikum/middleware"
	"praktikum/route"
)

func main() {
	config.InitDB()
	e := route.New()
	// implementasi middleware logger
	m.LogMiddleware(e)
	// start server
	e.Logger.Fatal(e.Start(":8000"))
}

package main

import (
	routes "github.com/naveenkakumanu/todo-echo/routes"
)

func main() {
	router := routes.Routes()
	router.Start(":8080")

}

package main

import (
	"app/router"
)

func main() {
	r := router.StartRouter()
	r.Run(":8080")
}

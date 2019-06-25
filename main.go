package main

import (
	"fmt"

	"./pkg/conf"
	"./router"
)

func main() {
	fmt.Println(conf.RunMode)

	r := router.InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}

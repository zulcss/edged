package main

import (
	"github.com/zulcss/edged/apiserver/pkg/controller"
)

func main() {
	r := controller.Setup()
	r.Run() // Listen and server on 0.0.0.0:8080
}

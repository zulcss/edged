package main

import (
	"github.com/zulcss/edged/apiserver/pkg/controller"
	"github.com/zulcss/edged/shared/db"
)

func main() {
	db.InitDatabase()

	r := controller.Setup()
	r.Run() // Listen and server on 0.0.0.0:8080
}

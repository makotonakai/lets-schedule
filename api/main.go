package main

import (
	"github.com/MakotoNakai/lets-schedule/api/router"
)

func main() {
    router := newRouter()
    router.Logger.Fatal(router.Start(":3000"))
}
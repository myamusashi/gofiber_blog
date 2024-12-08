package main

import (
	"log"
	"github.com/myamusashi/go_blog/internal/server"
)

func main()  {
    err := server.Start()
    if err != nil {
        log.Fatal(err)
    }
}

package main

import (
	"fmt"
	"github.com/eris/redis-example/routes"
	"log"
	"net/http"
)

func main()  {
	const port string = "3000"
	fmt.Printf("Api running on port %s\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
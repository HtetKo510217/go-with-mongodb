package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/htetko/go-with-mongodb/initializers"
	"github.com/htetko/go-with-mongodb/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("starting the applition ..")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("server is running on port 3000")

}

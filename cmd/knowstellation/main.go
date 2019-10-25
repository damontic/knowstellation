package main

import (
	"os"

	"github.com/damontic/knowstellation/pkg/client"
	"github.com/damontic/knowstellation/pkg/server"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "serve":
		server.Serve(args)
	case "server":
		server.Serve(args)
	default:
		endpoint := args[0]
		client.NewClient(endpoint)
	}
	/*
		knowstellationEndpoint := flag.String("knowstellation", "localhost:8080", "Knowstellationd endpoint")
		port := flag.String("port", ":8080", "Port to use")
		useHouston := flag.Bool("houston", false, "Decides to run the houston frontend")
		flag.Parse()

		endpoint, err := url.Parse(*knowstellationEndpoint)
		if err != nil {

		}
		client := client.NewClient(endpoint)
		client.Execute()

		user := core.NewUser("David Alberto Montaño Fetecua")
		golangStar := core.NewStar("GoLang")

		fmt.Println(user)
		fmt.Println(golangStar)
	*/
}

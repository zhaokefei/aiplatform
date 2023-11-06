package main

import "github.com/zhaokefei/aiplatform/api"


func main() {
	r := api.Routers()

	r.Run() // listen and serve on 0.0.0.0:8080
}

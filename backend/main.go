package main

import (
	"github.com/igroks/sd-project/backend/routes"
)

func main(){
	routes.HandleResquest()
	log.Fatal(http.ListenAndServe(":8000", nil))
}

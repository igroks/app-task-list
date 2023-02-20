package routes

import (
	"log"
	"net/http"
	"github.com/igroks/sd-project/backend/controllers"
)

func HandleResquest(){
	http.HandleFunc("/", controllers.Home)
}

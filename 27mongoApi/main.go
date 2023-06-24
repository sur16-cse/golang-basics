package main

import (
	"fmt"
	"net/http"

	"github.com/surbhikumari/mongoapi/router"
)

func main(){
    fmt.Println("MongoDBApi")
	fmt.Println("Server is getting started")
	r:=router.Router()
	http.ListenAndServe(":4000",r)
	fmt.Println("Listening at Port 4000 ...")
}
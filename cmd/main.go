package main

import (
	"fmt"
	"log"
	"net/http"

	"template/config"
	"template/service/server"

	"github.com/gin-gonic/gin"
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("error in loading LoadConfig()")
		log.Fatal(err.Error())
	}
}

func main() {
	r := gin.New()
	log.Fatal(http.ListenAndServe(":5002", server.NewServerImpl(r)))
}

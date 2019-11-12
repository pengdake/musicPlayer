package main

import (
	"github.com/gin-gonic/gin"

)


func main(){
	router := gin.Default()

	router.GET("/", index)
	router.POST("/track/play/:trackId", playTrack)
	router.POST("/track/stop", stopStrack)
	router.DELETE("/song/:songId", songDelete)

	router.Run(":8080")

}



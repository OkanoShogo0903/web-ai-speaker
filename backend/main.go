package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/OkanoShogo0903/web-ai-speaker/backend/controller"
	"github.com/OkanoShogo0903/web-ai-speaker/backend/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	speech_result := model.New()
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/speech", controller.SpeechPost(speech_result))

	router.Run(fmt.Sprintf(":%d", port))
}

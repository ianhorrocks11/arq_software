package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	Name string `json: "name`
}

func main() {
	engine:=gin.New()
	engine.POST("/test", func(context *gin.Context) { //funcion anonima 
		body:=Body{}
		//using BindJson method to serialize body with struct
		if err:=context.BindJSON(&body);err!=nil{
			context.AbortWithError(http.StatusBadRequest,err)
			return
		}
		fmt.Println(body)
		context.JSON(http.StatusAccepted,&body)
	})
	engine.Run(":3000")
}
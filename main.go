package main

import (
	"log"
	"login-app/configuration"
	"login-app/driver"
	"login-app/handlers"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

var (
	dynamoClient *dynamodb.Client
	logs         configuration.GoAppTools
)

func main() {
	InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

	logs.InfoLogger = *InfoLogger
	logs.ErrorLogger = *ErrorLogger
	var err error
	// chamada de função para a criação da sessao de login com o banco
	dynamoClient, err = driver.ConfigAws()
	//chamada da função para revificar o erro retornado
	configuration.Check(err, logs)

	appRouter := gin.New()
	appRouter.GET("/", func(ctx *gin.Context) {
		logs.InfoLogger.Println("Servidor Ok")
		handlers.ResponseOK(ctx, logs)
	})

	appRouter.GET("/logon", func(ctx *gin.Context) {
		handlers.GetUser(ctx, dynamoClient, logs)
	})

	appRouter.POST("/signin", func(ctx *gin.Context) {
		handlers.PostUser(ctx, dynamoClient, logs)
	})

	err = appRouter.Run()
	configuration.Check(err, logs)
	logs.InfoLogger.Println("Servidor terminou de iniciar.")
}

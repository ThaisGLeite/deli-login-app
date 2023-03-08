package handlers

import (
	"login-app/configuration"
	"login-app/database/query"
	"login-app/encrypt"
	"login-app/model"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, dynamoClient *dynamodb.Client, logs configuration.GoAppTools) {
	var userModel model.User
	err := c.BindJSON(&userModel)
	configuration.Check(err, logs)
	user := query.SelectUser(userModel.Nome, userModel.Senha, *dynamoClient, logs)
	if user.Nome == "" {
		c.IndentedJSON(http.StatusNotFound, "Nome de usuário "+userModel.Nome+" não encontrado")
		return
	}
	senhaTemp := encrypt.EncrytpHash(userModel.Senha, logs)
	if user.Senha != senhaTemp {
		logs.InfoLogger.Println(user.Senha)
		logs.InfoLogger.Println(senhaTemp)
		c.IndentedJSON(http.StatusUnauthorized, "Senha do usuário "+userModel.Nome+" não confere")
		return
	}
	//ToDO o usuário tem que ficar com a senha sempre do mesmo tamanho e tem que criptografar no inicio, nao no final so
}

func ResponseOK(c *gin.Context, app configuration.GoAppTools) {
	c.IndentedJSON(http.StatusOK, "Servidor up")
}

func PostUser(c *gin.Context, dynamoClient *dynamodb.Client, logs configuration.GoAppTools) {
	var newUser model.User

	//configue o model punh with the retorn of context gin
	err := c.BindJSON(&newUser)
	//faz a chacagem de errode forma unificada
	configuration.Check(err, logs)
	//call the package to encrypt password
	newUser.Senha = encrypt.EncrytpHash(newUser.Senha)
	//calling the quiry package to mount the request for a DB
	query.InsertUser(dynamoClient, newUser, logs)
	name := ("Colaborador " + newUser.Nome + " criado com sucesso!")
	c.IndentedJSON(http.StatusCreated, (name))
}

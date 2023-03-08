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

func MarshalUser(c *gin.Context, dynamoClient *dynamodb.Client, logs configuration.GoAppTools) (model.User, model.User) {
	var userModel model.User
	err := c.BindJSON(&userModel)
	configuration.Check(err, logs)
	user := query.SelectUser(userModel.Nome, userModel.Senha, *dynamoClient, logs)
	return userModel, user
}

// this funcion going into AWS, than user and passaword not is hash
func GetUser(c *gin.Context, dynamoClient *dynamodb.Client, logs configuration.GoAppTools) {
	userModel, user := MarshalUser(c, dynamoClient, logs)
	if user.Nome == "" {
		c.IndentedJSON(http.StatusNotFound, "Nome de usuário "+userModel.Nome+" não encontrado")
		return
	}
	if encrypt.CheckHash(userModel.Senha, user.Senha, logs) {
		c.IndentedJSON(http.StatusAccepted, "Authorized")
	} else {
		c.IndentedJSON(http.StatusUnauthorized, "Senha do usuário "+userModel.Nome+" não confere")
		return
	}
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
	newUser.Senha = encrypt.EncrytpHash(newUser.Senha, logs)
	//calling the quiry package to mount the request for a DB
	query.InsertUser(dynamoClient, newUser, logs)
	name := ("Colaborador " + newUser.Nome + " criado com sucesso!")
	c.IndentedJSON(http.StatusCreated, (name))
}

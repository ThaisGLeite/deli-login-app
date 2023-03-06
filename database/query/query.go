package query

import (
	"context"
	"login-app/configuration"
	"login-app/model"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/beevik/ntp"
)

func InsertUser(dynamoClient *dynamodb.Client, user model.User, logs configuration.GoAppTools) {

	//o codigo esta indo no observatorio nacional pegar a data e hora
	datatemp, err := ntp.Time("a.st1.ntp.br")
	configuration.Check(err, logs)

	//formatando a data retornada do observatorio para a data no formato desejado (yy-mm-dd_hh:mm)
	tempY := datatemp.Format("06")
	tempM := datatemp.Format("01")
	tempD := datatemp.Format("02")
	tempH := strconv.Itoa(datatemp.Hour())
	tempMin := strconv.Itoa(datatemp.Minute())
	user.DataCriacao = tempY + "-" + tempM + "-" + tempD + "_" + tempH + ":" + tempMin

	//converter a struct em um json
	item, err := attributevalue.MarshalMap(user)
	configuration.Check(err, logs)

	_, err = dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("LoginColaborador"),
		Item:      item,
	})
	configuration.Check(err, logs)
}

func SelectUser(Nome string, Senha string, dynamoClient dynamodb.Client, app configuration.GoAppTools) model.User {

	query := expression.Name("Nome").Equal(expression.Value(Nome))

	proj := expression.NamesList(expression.Name("Nome"), expression.Name("Senha"))

	expr, err := expression.NewBuilder().WithFilter(query).WithProjection(proj).Build()
	configuration.Check(err, app)

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("LoginColaborador"),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoClient.Scan(context.TODO(), params)
	configuration.Check(err, app)

	var user model.User
	for _, i := range result.Items {
		item := model.User{}

		err = attributevalue.UnmarshalMap(i, &item)

		configuration.Check(err, app)

		user = item
	}

	return user
}

package repository

import (
	"main/src/repository/entity"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var dynamo *dynamodb.DynamoDB

func CreateConnection() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})))
}

func PutItem(mutant entity.Mutants) error {

	dynamo = CreateConnection()

	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("mutants"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(mutant.Id),
			},
			"isMutant": {
				BOOL: aws.Bool(mutant.IsMutant),
			},
		},
	})

	return err
}

func GetItem(id string) (mutant entity.Mutants, err error) {

	dynamo = CreateConnection()

	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("mutants"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return mutant, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &mutant)

	return mutant, err
}

func GetCantItemsByIsMutant(isMutant bool) (int, error) {

	dynamo = CreateConnection()

	filt := expression.Name("isMutant").Equal(expression.Value(isMutant))

	proj := expression.NamesList(expression.Name("isMutant"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		return 0, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("mutants"),
	}

	result, err := dynamo.Scan(params)
	if err != nil {
		return 0, err
	}

	return int(*result.Count), nil
}

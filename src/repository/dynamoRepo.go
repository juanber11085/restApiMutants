package repository

import (
	"log"
	"main/src/repository/entity"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var dynamo *dynamodb.DynamoDB

//method used to create a connection to the DynamoDb database
func CreateConnection() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})))
}

//method used to insert a new record in the mutants table
func PutItem(mutant entity.Mutants) error {

	dynamo = CreateConnection()

	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("mutants"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(mutant.Id),
			},
			"isMutant": {
				N: aws.String(strconv.Itoa(int(mutant.IsMutant))),
			},
		},
	})

	return err
}

//method used to get a record by id
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

//method used to get the number of records by the isMutant field
func GetCantItemsByIsMutant(isMutant int8) (int, error) {

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

//method used to create the mutants table if it does not exist
func CreateTableIfNotExists() error {
	dynamo = CreateConnection()
	const table = "mutants"
	_, err := dynamo.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(table),
	})
	if awserr, ok := err.(awserr.Error); ok {
		if awserr.Code() == "ResourceNotFoundException" {
			_, err = dynamo.CreateTable(&dynamodb.CreateTableInput{
				TableName: aws.String(table),
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(int64(1)),
					WriteCapacityUnits: aws.Int64(int64(1)),
				},
				KeySchema: []*dynamodb.KeySchemaElement{{
					AttributeName: aws.String("id"),
					KeyType:       aws.String("HASH"),
				}, {
					AttributeName: aws.String("isMutant"),
					KeyType:       aws.String("RANGE"),
				}},
				AttributeDefinitions: []*dynamodb.AttributeDefinition{{
					AttributeName: aws.String("id"),
					AttributeType: aws.String("S"),
				}, {
					AttributeName: aws.String("isMutant"),
					AttributeType: aws.String("N"),
				}},
			})
			if err != nil {
				log.Print(err)
				return err
			}

			err = dynamo.WaitUntilTableExists(&dynamodb.DescribeTableInput{
				TableName: aws.String(table),
			})
			if err != nil {
				log.Print(err)
				return err
			}
		}
	}
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

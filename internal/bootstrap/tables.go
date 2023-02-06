package bootstrap

//import (
//	"fmt"

//	"github.com/Skavengerr/expone/util"
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/service/dynamodb"
//)

//func CreateTableExpenses(dynamo *dynamodb.DynamoDB) {
//	if dynamo == nil {
//		return
//	}

//	result, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
//		AttributeDefinitions: []*dynamodb.AttributeDefinition{
//			{AttributeName: aws.String("account_id"), AttributeType: aws.String("S")},
//		},
//		KeySchema: []*dynamodb.KeySchemaElement{
//			{AttributeName: aws.String("account_id"), KeyType: aws.String("HASH")},
//		},
//		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
//			ReadCapacityUnits:  aws.Int64(10),
//			WriteCapacityUnits: aws.Int64(10),
//		},
//		TableName: aws.String(repository.TABLE_EXPENSE),
//	})

//	fmt.Println("Created the table", result)

//	if err != nil {
//		util.HandleAWSError(err)
//		return
//	}
//}

//func CreateTableIncomes(dynamo *dynamodb.DynamoDB) {
//	if dynamo == nil {
//		return
//	}

//	result, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
//		AttributeDefinitions: []*dynamodb.AttributeDefinition{
//			{AttributeName: aws.String("account_id"), AttributeType: aws.String("S")},
//		},
//		KeySchema: []*dynamodb.KeySchemaElement{
//			{AttributeName: aws.String("account_id"), KeyType: aws.String("HASH")},
//		},
//		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
//			ReadCapacityUnits:  aws.Int64(10),
//			WriteCapacityUnits: aws.Int64(10),
//		},
//		TableName: aws.String(model.TABLE_INCOME),
//	})

//	fmt.Println("Created the table", result)

//	if err != nil {
//		util.HandleAWSError(err)
//		return
//	}
//}

//func CreateTableAccounts(dynamo *dynamodb.DynamoDB) {
//	if dynamo == nil {
//		return
//	}

//	result, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
//		AttributeDefinitions: []*dynamodb.AttributeDefinition{
//			{AttributeName: aws.String("account_id"), AttributeType: aws.String("S")},
//		},
//		KeySchema: []*dynamodb.KeySchemaElement{
//			{AttributeName: aws.String("account_id"), KeyType: aws.String("HASH")},
//		},
//		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
//			ReadCapacityUnits:  aws.Int64(10),
//			WriteCapacityUnits: aws.Int64(10),
//		},
//		TableName: aws.String(model.TABLE_ACCOUNTS),
//	})

//	fmt.Println("Created the table", result)

//	if err != nil {
//		util.HandleAWSError(err)
//		return
//	}
//}

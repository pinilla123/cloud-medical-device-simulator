package device

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var svc *dynamodb.Client // DynamoDB client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc = dynamodb.NewFromConfig(cfg)
}

func SaveVitalSigns(vs VitalSigns) error {
	av, err := attributevalue.MarshalMap(vs)
	if err != nil {
		log.Printf("Failed to marshal vital signs: %v", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("MedicalDeviceData"),
		Item:      av,
	}

	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to put item in DynamoDB: %v", err)
		return err
	}

	return nil
}

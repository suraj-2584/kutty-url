package utils

import (
	"deps/url-shortener/internal/constants"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"context"
)

func IsUrlValid(inputUrl string) bool {
	u, err := url.ParseRequestURI(inputUrl)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IntToBase62String(num int32) string {
	if num == 0 {
		return "0"
	}
	var encoded []byte

	for num > 0 {
		remainder := num % 62
		encoded = append([]byte{constants.Base62Chars[remainder]}, encoded...)
		num /= 62
	}
	for len(encoded) < 5 {
		encoded = append([]byte{constants.Base62Chars[0]}, encoded...)
	}
	return string(encoded)
}

func LoadSecrets() {
	secretName := "kutty-url"
	region := "ap-south-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	constants.ConnectionString = *result.SecretString
}

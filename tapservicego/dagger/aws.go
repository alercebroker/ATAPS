package main

import (
	"context"
	"dagger/tapservicego/internal/dagger"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func withAWSCredentials(container *dagger.Container) *dagger.Container {
	return container.
		WithEnvVariable("AWS_ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID")).
		WithEnvVariable("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY")).
		WithEnvVariable("AWS_SESSION_TOKEN", os.Getenv("AWS_SESSION_TOKEN"))
}

func getSsmValue(parameterName string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}
	// Create an SSM client
	client := ssm.NewFromConfig(cfg)
	parameter, err := client.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: &parameterName,
	})
	if err != nil {
		return "", err
	}
	return *parameter.Parameter.Value, nil
}

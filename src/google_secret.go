package main

import (
	"context"
	"fmt"

	"github.com/fatih/color"

	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
)

type GoogleSecret struct {
	Project    string `yaml:"project"`
	SecretName string `yaml:"name"`
}

func (g *GoogleSecret) toSecret() (Secret, error) {
	fmt.Println("hi!!!!!")

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)

	if err != nil {
		color.Red("Failed to setup context for Google Cloud.")
		return Secret{}, err
	}

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: g.SecretName,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)

	if err != nil {
		color.Red("Failed to grab secret from Google Cloud Secrets Manager:")
		color.Cyan(err.Error())

		return Secret{}, err
	}

	secret := Secret{
		Key:   g.SecretName,
		Value: string(result.Payload.Data),
	}

	return secret, nil
}

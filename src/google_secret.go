package main

import (
	"context"
	"fmt"

	"github.com/fatih/color"

	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
)

// GoogleSecret is a representation of a GoogleSecretManager request
type GoogleSecret struct {
	Project    string `yaml:"project"`
	SecretName string `yaml:"name"`
	Version    int    `yaml:"version"`
}

func (g *GoogleSecret) toSecret() (Secret, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)

	if err != nil {
		color.Red("Failed to setup context for Google Cloud.")
		color.Red(err.Error())
		return Secret{}, err
	}

	// construct the secrets path, found in doctstring for `secretmanagerpb.AccessSecretVersionRequest`
	path := fmt.Sprintf("projects/%s/secrets/%s/versions/%d", g.Project, g.SecretName, g.Version)
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: path,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)

	if err != nil {
		color.Red("Failed to grab secret from Google Cloud Secrets Manager:")
		color.Cyan("❌ " + err.Error())

		return Secret{}, err
	}

	secret := Secret{
		Key:   g.SecretName,
		Value: string(result.Payload.Data),
	}

	return secret, nil
}

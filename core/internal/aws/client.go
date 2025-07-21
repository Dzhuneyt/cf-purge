package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func NewCloudFormationClient() *cloudformation.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to load AWS configuration: %v", err)
	}

	// Check if AWS credentials are available
	if _, err = cfg.Credentials.Retrieve(context.TODO()); err != nil {
		log.Fatalf("Failed to retrieve AWS credentials: %v", err)
	}

	return cloudformation.NewFromConfig(cfg)
}

package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// TODO: Upload to github
// TODO: Write a build script that call go build and copies this program to a folder which is in $PATH
func main() {
	// Load the default AWS configuration from the environment (e.g., ~/.aws/config, ~/.aws/credentials)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}
	// Create an STS client
	stsClient := sts.NewFromConfig(cfg)
	// Call GetCallerIdentity to get information about the current IAM role/user
	output, err := stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("unable to get caller identity, %v", err)
	}
	boldStart := "\033[1m"
	boldEnd := "\033[0m"
	fmt.Printf("Account ID: %s%s%s\n", boldStart, aws.ToString(output.Account), boldEnd)
	fmt.Printf("Role:       %s\n", aws.ToString(&strings.Split(*output.Arn, "/")[1]))
	fmt.Printf("ARN:        %s\n", aws.ToString(output.Arn))
}

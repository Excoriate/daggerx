package cloudaws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// AWSAdapter defines an interface for creating various AWS service clients.
type AWSAdapter interface {
	// NewSNS creates a new Simple Notification Service (SNS) client.
	NewSNS() *sns.Client

	// NewSQS creates a new Simple Queue Service (SQS) client.
	NewSQS() *sqs.Client

	// NewS3 creates a new Simple Storage Service (S3) client.
	NewS3() *s3.Client

	// NewRDS creates a new Relational Database Service (RDS) client.
	NewRDS() *rds.Client

	// NewEC2 creates a new Elastic Compute Cloud (EC2) client.
	NewEC2() *ec2.Client

	// NewIAM creates a new Identity and Access Management (IAM) client.
	NewIAM() *iam.Client

	// NewDynamoDB creates a new DynamoDB client.
	NewDynamoDB() *dynamodb.Client

	// NewAutoScaling creates a new Auto Scaling client.
	NewAutoScaling() *autoscaling.Client

	// NewECS creates a new Elastic Container Service (ECS) client.
	NewECS() *ecs.Client

	// NewEKS creates a new Elastic Kubernetes Service (EKS) client.
	NewEKS() *eks.Client
}

// AWS implements the AWSAdapter interface and holds the configuration for AWS services.
type AWS struct {
	Region string
	cfg    aws.Config
}

// NewAWS creates a new instance of AWS with the specified region.
// It requires AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables to be set.
//
// Parameters:
//   - region: The AWS region to use.
//
// Returns:
//   - AWSAdapter: An interface for creating AWS service clients.
//   - error: An error if the AWS configuration could not be loaded.
func NewAWS(region string) (AWSAdapter, error) {
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		return nil, fmt.Errorf("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set")
	}

	cfg, err := awscfg.LoadDefaultConfig(context.TODO(),
		awscfg.WithRegion(region),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	return &AWS{Region: region, cfg: cfg}, nil
}

// NewSNS creates a new Simple Notification Service (SNS) client.
//
// Returns:
//   - *sns.Client: A new SNS client.
func (a *AWS) NewSNS() *sns.Client {
	return sns.NewFromConfig(a.cfg)
}

// NewSQS creates a new Simple Queue Service (SQS) client.
//
// Returns:
//   - *sqs.Client: A new SQS client.
func (a *AWS) NewSQS() *sqs.Client {
	return sqs.NewFromConfig(a.cfg)
}

// NewS3 creates a new Simple Storage Service (S3) client.
//
// Returns:
//   - *s3.Client: A new S3 client.
func (a *AWS) NewS3() *s3.Client {
	return s3.NewFromConfig(a.cfg)
}

// NewRDS creates a new Relational Database Service (RDS) client.
//
// Returns:
//   - *rds.Client: A new RDS client.
func (a *AWS) NewRDS() *rds.Client {
	return rds.NewFromConfig(a.cfg)
}

// NewEC2 creates a new Elastic Compute Cloud (EC2) client.
//
// Returns:
//   - *ec2.Client: A new EC2 client.
func (a *AWS) NewEC2() *ec2.Client {
	return ec2.NewFromConfig(a.cfg)
}

// NewIAM creates a new Identity and Access Management (IAM) client.
//
// Returns:
//   - *iam.Client: A new IAM client.
func (a *AWS) NewIAM() *iam.Client {
	return iam.NewFromConfig(a.cfg)
}

// NewDynamoDB creates a new DynamoDB client.
//
// Returns:
//   - *dynamodb.Client: A new DynamoDB client.
func (a *AWS) NewDynamoDB() *dynamodb.Client {
	return dynamodb.NewFromConfig(a.cfg)
}

// NewAutoScaling creates a new Auto Scaling client.
//
// Returns:
//   - *autoscaling.Client: A new Auto Scaling client.
func (a *AWS) NewAutoScaling() *autoscaling.Client {
	return autoscaling.NewFromConfig(a.cfg)
}

// NewECS creates a new Elastic Container Service (ECS) client.
//
// Returns:
//   - *ecs.Client: A new ECS client.
func (a *AWS) NewECS() *ecs.Client {
	return ecs.NewFromConfig(a.cfg)
}

// NewEKS creates a new Elastic Kubernetes Service (EKS) client.
//
// Returns:
//   - *eks.Client: A new EKS client.
func (a *AWS) NewEKS() *eks.Client {
	return eks.NewFromConfig(a.cfg)
}

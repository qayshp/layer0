package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type Client struct {
	AutoScaling    autoscalingiface.AutoScalingAPI
	CloudTrail     cloudtrailiface.CloudTrailAPI
	CloudWatchLogs cloudwatchlogsiface.CloudWatchLogsAPI
	EC2            ec2iface.EC2API
	ECS            ecsiface.ECSAPI
	ELB            elbiface.ELBAPI
	IAM            iamiface.IAMAPI
	S3             s3iface.S3API
}

func NewClient(session *session.Session) *Client {
	return &Client{
		AutoScaling:    autoscaling.New(session),
		CloudTrail:     cloudtrail.New(session),
		CloudWatchLogs: cloudwatchlogs.New(session),
		EC2:            ec2.New(session),
		ECS:            ecs.New(session),
		ELB:            elb.New(session),
		IAM:            iam.New(session),
		S3:             s3.New(session),
	}
}
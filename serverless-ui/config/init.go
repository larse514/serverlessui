package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/larse514/aws-cloudformation-go"
	"github.com/larse514/serverlessui/serverless-ui/actions"
	"github.com/larse514/serverlessui/serverless-ui/bucket"
	"github.com/larse514/serverlessui/serverless-ui/commands"
	"github.com/larse514/serverlessui/serverless-ui/dns"
	"github.com/larse514/serverlessui/serverless-ui/fileutil"
	"github.com/larse514/serverlessui/serverless-ui/iaas"
	"github.com/urfave/cli"
)

//CreateApp method to create initial app
func CreateApp() *cli.App {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		log.Fatal("error creating session")
		os.Exit(1)
	}

	cloudformation := cloudformation.New(sess)
	s3 := s3.New(sess)
	resource := cf.Stack{Client: cloudformation}
	executor := cf.IaaSExecutor{Client: cloudformation}
	dns := dns.Route53{Executor: &executor, Resource: resource}
	s3Bucket := bucket.S3Bucket{Executor: executor, Resource: resource, IaaS: iaas.AWSTemplate{}}
	uploader := bucket.S3Uploader{Client: s3, FileUtil: fileutil.FileUtility{}}
	deployAction := actions.ServerlessUI{DNS: dns, Bucket: s3Bucket, Uploader: uploader}
	app := cli.NewApp()

	app.Name = "serverless-ui"
	app.Usage = "Command line interface for serverless ui deployment"
	app.Version = "0.0.1"
	app.Author = "VSS"
	app.Commands = []cli.Command{
		commands.Deploy(deployAction),
	}

	return app
}

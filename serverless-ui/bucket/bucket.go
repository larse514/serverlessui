package bucket

import (
	"log"
	"strings"

	"github.com/larse514/aws-cloudformation-go"
	"github.com/larse514/serverlessui/serverless-ui/commands"
)

const (
	s3BucketPath = "https://s3.amazonaws.com/serverless-ui-deployables/s3site.yml"
	//route53 param values
	domainNameParam     = "HostedZone"
	fullDomainNameParam = "FullDomainName"
	acmCertARNParam     = "AcmCertificateArn"
	ttlCacheValueParam  = "CacheValueTTL"
)

//S3Bucket is a struct to define needed S3 Bucket dependencies
type S3Bucket struct {
	Executor cf.Executor
	Resource cf.Resource
}

//DeploySite is a function to Create an S3 Site with CDN and ACM
func (s3Bucket S3Bucket) DeploySite(input *commands.BucketInput) error {
	stackName := getStackName(input)

	stack, err := s3Bucket.Resource.GetStack(&stackName)
	if err != nil {
		return err
	}
	if *stack.StackName == "" {
		log.Println("Creating s3 bucket ", stack)
		//create stack
		err = s3Bucket.Executor.CreateStackFromS3(s3BucketPath, stackName, createInputParameters(input), nil)
		if err != nil {
			return err
		}
		return s3Bucket.Executor.PauseUntilCreateFinished(stackName)
	}
	log.Println("S3 bucket already exists")
	return nil

}

//Method to convert DomainName from input to stack name
//route53 does not allow for full stop (.) characters
func getStackName(input *commands.BucketInput) string {
	return strings.Replace(input.FullDomainName, ".", "-", -1)
}

//Helper method to create []*cloudformation.Parameter from input
func createInputParameters(input *commands.BucketInput) *map[string]string {
	//we need to convert this (albeit awkwardly for the time being) to Cloudformation Parameters
	//we do as such first by converting everything to a key value map
	//key being the CF Param name, value is the value to provide to the cloudformation template
	parameterMap := make(map[string]string, 0)
	//todo-refactor this bloody hardcoded mess
	parameterMap[domainNameParam] = input.HostedZone
	parameterMap[fullDomainNameParam] = input.FullDomainName
	parameterMap[acmCertARNParam] = input.AcmCertificateArn
	parameterMap[ttlCacheValueParam] = input.CacheValueTTL

	return &parameterMap

}

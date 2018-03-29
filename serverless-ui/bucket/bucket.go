package bucket

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/larse514/serverlessui/serverless-cli/cf"
)

const (
	s3BucketPath = "ias/cloudformation/s3site.yml"
	//route53 param values
	domainNameParam     = "DomainName"
	fullDomainNameParam = "FullDomainName"
	acmCertARNParam     = "AcmCertificateArn"
	ttlCacheValueParam  = "CacheValueTTL"
)

//Bucket is an interface to define creation of Bucket based sites
type Bucket interface {
	CreateSite(input *Input) error
}

//S3Bucket is a struct to define needed S3 Bucket dependencies
type S3Bucket struct {
	Executor cf.Executor
}

//Input is a struct which deines the required parameters to create an s3 bucket based site
type Input struct {
	DomainName        string
	FullDomainName    string
	AcmCertificateArn string
	CacheValueTTL     string
}

//CreateSite is a function to Create an S3 Site with CDN and ACM
func (s3Bucket S3Bucket) CreateSite(input *Input) error {
	stackName := getStackName(input)
	//create stack
	s3Bucket.Executor.CreateStack(s3BucketPath, stackName, createInputParameters(input))

	return s3Bucket.Executor.PauseUntilCreateFinished(stackName)
}

//Method to convert DomainName from input to stack name
//route53 does not allow for full stop (.) characters
func getStackName(input *Input) string {
	return strings.Replace(input.FullDomainName, ".", "-", -1)
}

//Helper method to create []*cloudformation.Parameter from input
func createInputParameters(input *Input) []*cloudformation.Parameter {
	//we need to convert this (albeit awkwardly for the time being) to Cloudformation Parameters
	//we do as such first by converting everything to a key value map
	//key being the CF Param name, value is the value to provide to the cloudformation template
	parameterMap := make(map[string]string, 0)
	//todo-refactor this bloody hardcoded mess
	parameterMap[domainNameParam] = input.DomainName
	parameterMap[fullDomainNameParam] = input.FullDomainName
	parameterMap[acmCertARNParam] = input.AcmCertificateArn
	parameterMap[ttlCacheValueParam] = input.CacheValueTTL

	return cf.CreateCloudformationParameters(parameterMap)

}

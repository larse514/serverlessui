package dns

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/larse514/serverlessui/serverless-cli/cf"
)

const (
	route53Path = "ias/cloudformation/route53.yml"
	//route53 param values
	domainNameParam       = "DomainName"
	hostedZoneExistsParam = "HostedZoneExists"
	environmentParam      = "Environment"
)

//DNS is an interface to represent Cloud DNS Services
type DNS interface {
	CreateHostedZone(input *Input) error
}

//Route53 is an implementation of the DNS interface
type Route53 struct {
	Executor cf.Executor
}

//Input is a struct representing the required parameters to pass for HostedZoneCreation creation
type Input struct {
	DomainName string
	//todo- add type safety
	HostedZoneExists string
	Environment      string
}

//CreateHostedZone Method to create Route53 hosted zone
func (route53 Route53) CreateHostedZone(input *Input) error {
	//replace domain name
	stackName := getStackName(input)
	//todo-check for existance
	//create stack
	route53.Executor.CreateStack(route53Path, stackName, createDNSInputParameters(input))

	return route53.Executor.PauseUntilCreateFinished(stackName)

}

//Method to convert DomainName from input to stack name
//route53 does not allow for full stop (.) characters
func getStackName(input *Input) string {
	return strings.Replace(input.DomainName, ".", "-", -1)
}

//Helper method to create []*cloudformation.Parameter from input
func createDNSInputParameters(input *Input) []*cloudformation.Parameter {
	//we need to convert this (albeit awkwardly for the time being) to Cloudformation Parameters
	//we do as such first by converting everything to a key value map
	//key being the CF Param name, value is the value to provide to the cloudformation template
	parameterMap := make(map[string]string, 0)
	//todo-refactor this bloody hardcoded mess
	parameterMap[domainNameParam] = input.DomainName
	parameterMap[environmentParam] = input.Environment
	parameterMap[hostedZoneExistsParam] = input.HostedZoneExists

	return cf.CreateCloudformationParameters(parameterMap)

}

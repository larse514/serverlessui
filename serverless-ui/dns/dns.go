package dns

import (
	"log"
	"strings"

	"github.com/larse514/aws-cloudformation-go"
	"github.com/larse514/serverlessui/serverless-ui/commands"
)

const (
	route53Path = "https://s3.amazonaws.com/serverless-ui-deployables/route53.yml"
	//route53 param values
	domainNameParam       = "HostedZone"
	hostedZoneExistsParam = "HostedZoneExists"
	environmentParam      = "Environment"
)

//Route53 is an implementation of the DNS interface
type Route53 struct {
	Executor cf.Executor
	Resource cf.Resource
}

//DeployHostedZone Method to create Route53 hosted zone
func (route53 Route53) DeployHostedZone(input *commands.DNSInput) error {
	//replace domain name
	log.Println(*input)
	stackName := getStackName(input)
	log.Println(stackName)

	//todo- i recommend refactoring this out of here
	stack, err := route53.Resource.GetStack(&stackName)
	if err != nil {
		return err
	}
	if stack.StackName == nil {
		//create stack
		route53.Executor.CreateStackFromS3(route53Path, stackName, createDNSInputParameters(input), nil)
		return route53.Executor.PauseUntilCreateFinished(stackName)
	}

	return nil
}

//Method to convert DomainName from input to stack name
//route53 does not allow for full stop (.) characters
func getStackName(input *commands.DNSInput) string {
	return strings.Replace(input.HostedZone, ".", "-", -1)
}

//Helper method to create []*cloudformation.Parameter from input
func createDNSInputParameters(input *commands.DNSInput) *map[string]string {
	//we need to convert this (albeit awkwardly for the time being) to Cloudformation Parameters
	//we do as such first by converting everything to a key value map
	//key being the CF Param name, value is the value to provide to the cloudformation template
	parameterMap := make(map[string]string, 0)
	//todo-refactor this bloody hardcoded mess
	parameterMap[domainNameParam] = input.HostedZone
	parameterMap[environmentParam] = input.Environment
	parameterMap[hostedZoneExistsParam] = input.HostedZoneExists

	return &parameterMap

}

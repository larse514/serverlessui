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
	websiteArnOutput      = "WebsiteCertArn"
)

//Route53 is an implementation of the DNS interface
type Route53 struct {
	Executor cf.Executor
	Resource cf.Resource
}

//Route53Output struct containing output from Route53
type Route53Output struct {
	WebsiteArn string
}

//DeployHostedZone Method to create Route53 hosted zone
func (route53 Route53) DeployHostedZone(input *commands.DNSInput) (*Route53Output, error) {
	//replace domain name
	stackName := getStackName(input)

	//todo- i recommend refactoring this out of here
	stack, err := route53.Resource.GetStack(&stackName)
	if err != nil {
		return nil, err
	}
	websiteOutputValue := input.Environment + "-" + websiteArnOutput
	if *stack.StackName == "" {
		log.Println("Creating new dns stack")
		//create stack
		err = route53.Executor.CreateStackFromS3(route53Path, stackName, createDNSInputParameters(input), nil)
		if err != nil {
			return nil, err
		}
		err = route53.Executor.PauseUntilCreateFinished(stackName)
		if err != nil {
			return nil, err
		}
		stack, err = route53.Resource.GetStack(&stackName)
		if err != nil {
			return nil, err
		}
		return &Route53Output{WebsiteArn: cf.GetOutputValue(stack, websiteOutputValue)}, nil

	}

	log.Println("DNS Stack already exists")
	return &Route53Output{WebsiteArn: cf.GetOutputValue(stack, websiteOutputValue)}, nil
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

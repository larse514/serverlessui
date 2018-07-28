package bucket

import (
	"testing"

	"github.com/larse514/serverlessui/serverless-ui/commands"
)

const (
	domainName       = "somedomain.com"
	longerDomainName = "prefix.somedomain.com"
)

func TestGetStackName(t *testing.T) {
	input := commands.BucketInput{FullDomainName: domainName}
	expected := "somedomain-com"
	got := getStackName(&input)

	if expected != got {
		t.Log("Received ", got, " expected ", expected)
		t.Fail()
	}
}

func TestGetStackNamePrefix(t *testing.T) {
	input := commands.BucketInput{FullDomainName: longerDomainName}

	expected := "prefix-somedomain-com"
	got := getStackName(&input)

	if expected != got {
		t.Log("Received ", got, " expected ", expected)
		t.Fail()
	}
}

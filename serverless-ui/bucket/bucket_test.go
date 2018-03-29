package bucket

import "testing"

const (
	domainName       = "somedomain.com"
	longerDomainName = "prefix.somedomain.com"
)

func TestGetStackName(t *testing.T) {
	input := Input{FullDomainName: domainName}
	expected := "somedomain-com"
	got := getStackName(&input)

	if expected != got {
		t.Log("Received ", got, " expected ", expected)
		t.Fail()
	}
}

func TestGetStackNamePrefix(t *testing.T) {
	input := Input{FullDomainName: longerDomainName}

	expected := "prefix-somedomain-com"
	got := getStackName(&input)

	if expected != got {
		t.Log("Received ", got, " expected ", expected)
		t.Fail()
	}
}

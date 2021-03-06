package commands

import (
	"github.com/larse514/serverlessui/serverless-ui/flags"
	"github.com/urfave/cli"
)

const (
	hostedZone          = "hostedzone"
	domainName          = "domainname"
	cacheTTLArg         = "cachettl"
	hostedZoneExistsArg = "hostedzoneexists"
	tagArg              = "tag"
	environment         = "environment"
	appDir              = "applicationdirectory"
)

//DNSInput is a struct representing the required parameters to pass for HostedZoneCreation creation
type DNSInput struct {
	HostedZone string
	//todo- add type safety
	HostedZoneExists string
	Environment      string
}

//BucketInput is a struct which deines the required parameters to create an s3 bucket based site
type BucketInput struct {
	HostedZone        string
	FullDomainName    string
	AcmCertificateArn string
	CacheValueTTL     string
}

//DeployAction interface to define deploy action
type DeployAction interface {
	Deploy(dnsInput *DNSInput, bucketInput *BucketInput, appDir string) error
}

//Deploy is a command to deploy the UI
func Deploy(action DeployAction) cli.Command {
	return cli.Command{

		Name:    "deploy",
		Aliases: []string{"d"},
		Usage:   "Deploy ui application",
		Flags:   flags.Deploy(),
		Action: func(c *cli.Context) error {
			dnsInput := DNSInput{
				HostedZone:       c.String(hostedZone),
				HostedZoneExists: c.String(hostedZoneExistsArg),
				Environment:      c.String(environment),
			}
			bucketinput := BucketInput{
				HostedZone:     c.String(hostedZone),
				FullDomainName: c.String(domainName),
				CacheValueTTL:  c.String(cacheTTLArg),
			}
			return action.Deploy(&dnsInput, &bucketinput, c.String(appDir))
		},
	}
}
